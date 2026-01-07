# AIChaseComponent

플레이어나 다른 엔티티를 추적하는 AI입니다. StateComponent가 없다면 자동으로 추가됩니다.

# Properties

| float DetectionRange |
| --- |
| 추적 감지 거리입니다. 대상의 거리가 이보다 멀어지면 추적을 중단하고, 가까워지면 다시 추적을 시작합니다. |

| boolean IsChaseNearPlayer |
| --- |
| 값이 true인 경우 DetectionRange 프로퍼티 값 내에서 가장 가까운 플레이어를 자동으로 추적합니다. TargetEntityRef 프로퍼티 또는 SetTarget(Entity) 함수로 지정한 대상이 있는 경우 플레이어 대신 지정한 대상을 추적합니다. |

| [EntityRef](https://mod-developers.nexon.com/apiReference/Misc/EntityRef) TargetEntityRef ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 추적 대상이 될 엔티티를 정합니다. |

##### inherited from AIComponent:

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Legacy 시스템 지원 여부를 설정합니다. 이전 시스템은 노드의 Running 상태를 구현하기 위해 ExclusiveExecutionWhenRunning 프로퍼티를 사용해야 합니다. Legacy 시스템은 더 이상 지원하지 않으며, 추후 삭제 예정입니다. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 값이 true면 행동 트리가 실행될 때 실행 관련 정보가 로그로 출력됩니다. 메이커 환경에서만 동작합니다. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 업데이트 권한입니다. Server와 Client 중 설정한 곳에서 실행됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetCurrentTarget() |
| --- |
| 추적 대상 엔티티를 반환합니다. |

| void SetTarget([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity) |
| --- |
| targetEntity를 추적하도록 설정합니다. 함수 호출 시 IsChaseNearPlayer 프로퍼티는 자동으로 비활성화됩니다 |

##### inherited from AIComponent:

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Action 노드를 생성합니다. 노드가 실행될 때 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| BTNodeType을 기반으로 Action 노드를 생성합니다. nodeType은 BTNodeType의 타입명입니다. onBehaveFunction가 nil이 아니라면 노드가 실행될 때 BTNodeType의 OnInit(), OnBehave() 함수 대신 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| node를 최상위 노드로 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

몬스터가 가까이 있는 적을 무작정 따라가지 않고, 자신을 공격한 대상을 추적하는 예제입니다. Chase 몬스터에 예제 코드를 작성한 컴포넌트를 추가해 사용할 수 있습니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	local aiChaseComponent = self.Entity.AIChaseComponent
	if aiChaseComponent == nil then
		return
	end
	 
	aiChaseComponent.IsChaseNearPlayer = false
	 
	local chatBallon = self.Entity.ChatBalloonComponent
	if chatBallon == nil then
		chatBallon = self.Entity:AddComponent(ChatBalloonComponent)
	end
	 
	chatBallon.ChatModeEnabled = false
	chatBallon.ShowDuration = 1
	chatBallon.HideDuration = 0
	chatBallon.FontSize = 1.2
}

[server only]
void OnUpdate (number delta)
{
	if self.Entity.ChatBalloonComponent == nil then
		return
	end

	local currentTargetEntity = self.Entity.AIChaseComponent:GetCurrentTarget()
	if currentTargetEntity == nil then
		self.Entity.ChatBalloonComponent.AutoShowEnabled = false
	else
		self.Entity.ChatBalloonComponent.AutoShowEnabled = true
		self.Entity.ChatBalloonComponent.Message = "target is "..currentTargetEntity.Name
	end
}

Event Handler:
[self]
HandleHitEvent (HitEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: HitComponent
	-- Space: Server, Client
	---------------------------------------------------------
	 
	-- Parameters
	local AttackCenter = event.AttackCenter
	local AttackerEntity = event.AttackerEntity
	local Damages = event.Damages
	local Extra = event.Extra
	local FeedbackAction = event.FeedbackAction
	local IsCritical = event.IsCritical
	local TotalDamage = event.TotalDamage
	---------------------------------------------------------
	if self.Entity.AIChaseComponent == nil then
		return
	end
	 
	self.Entity.AIChaseComponent:SetTarget(AttackerEntity)
}
```

# SeeAlso

- [AIComponent](https://mod-developers.nexon.com/apiReference/Components/AIComponent)
- [StateComponent](https://mod-developers.nexon.com/apiReference/Components/StateComponent)

Update 2025-08-27 PM 04:56


# AIComponent

엔티티에 AI를 부여하기 위한 행동 트리 기능을 제공합니다.

# Properties

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Legacy 시스템 지원 여부를 설정합니다. 이전 시스템은 노드의 Running 상태를 구현하기 위해 ExclusiveExecutionWhenRunning 프로퍼티를 사용해야 합니다. Legacy 시스템은 더 이상 지원하지 않으며, 추후 삭제 예정입니다. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 값이 true면 행동 트리가 실행될 때 실행 관련 정보가 로그로 출력됩니다. 메이커 환경에서만 동작합니다. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 업데이트 권한입니다. Server와 Client 중 설정한 곳에서 실행됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Action 노드를 생성합니다. 노드가 실행될 때 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| BTNodeType을 기반으로 Action 노드를 생성합니다. nodeType은 BTNodeType의 타입명입니다. onBehaveFunction가 nil이 아니라면 노드가 실행될 때 BTNodeType의 OnInit(), OnBehave() 함수 대신 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| node를 최상위 노드로 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

몬스터가 잠을 자다가 플레이어가 다가오면 다가오지 말라고 말하는 예제입니다. AIComponent를 Extend해 사용합니다. 작성한 스크립트 컴포넌트를 static 몬스터에 추가해 사용합니다.

```
Property:
[Sync]
number DetectDistance = 4

Method:
[server only]
void OnBeginPlay ()
{
	local chatBallon = self.Entity.ChatBalloonComponent
	if chatBallon == nil then
		chatBallon = self.Entity:AddComponent(ChatBalloonComponent)
	end
	 
	self._T.nearestPlayer = nil
	 
	chatBallon.AutoShowEnabled = true
	chatBallon.ChatModeEnabled = false
	chatBallon.ShowDuration = 1
	chatBallon.HideDuration = 0
	chatBallon.FontSize = 1.5
	 
	local function isNearPlayer(deltaTime)
		local players = _UserService:GetUsersByMapComponent(self.Entity.CurrentMap.MapComponent)
		self._T.nearestPlayer = nil
		local dist = math.maxinteger
		for i, player in pairs(players) do
			if isvalid(player) then
				local distTemp = Vector2.Distance(
				player.TransformComponent.Position:ToVector2(), self.Entity.TransformComponent.Position:ToVector2())
				dist = math.min(dist, distTemp)
				if dist <= self.DetectDistance then
					self._T.nearestPlayer = player
				end
			end
		end
	     
		if self._T.nearestPlayer == nil then
			return BehaviourTreeStatus.Failure
		else
			return BehaviourTreeStatus.Success
		end
	end
	 
	local function lookAtNearestPlayer(deltaTime)
		local flipX = self.Entity.TransformComponent.Position.x < self._T.nearestPlayer.TransformComponent.Position.x
		self.Entity.SpriteRendererComponent.FlipX = flipX
		return BehaviourTreeStatus.Success
	end
	 
	local function warn(deltaTime)
		chatBallon.Message = "Don't come!"
		return BehaviourTreeStatus.Success
	end
	 
	local function sleep(deltaTime)
		chatBallon.Message = "Zzz..."
		return BehaviourTreeStatus.Success
	end
	 
	local rootNode = SelectorNode("Root")
	 
	local alertSeq = SequenceNode("AlertSequence")
	alertSeq:AttachChild(self:CreateLeafNode("IsNearPlayer", isNearPlayer))
	alertSeq:AttachChild(self:CreateLeafNode("LookAtNearestPlayer", lookAtNearestPlayer))
	alertSeq:AttachChild(self:CreateLeafNode("Warn", warn))
	 
	rootNode:AttachChild(alertSeq)
	rootNode:AttachChild(self:CreateLeafNode("Sleep", sleep))
	 
	self:SetRootNode(rootNode)
}
```

# SeeAlso

- [ChatBalloonComponent](https://mod-developers.nexon.com/apiReference/Components/ChatBalloonComponent)
- [SpriteRendererComponent](https://mod-developers.nexon.com/apiReference/Components/SpriteRendererComponent)
- [TransformComponent](https://mod-developers.nexon.com/apiReference/Components/TransformComponent)
- [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode)
- [SelectorNode](https://mod-developers.nexon.com/apiReference/Misc/SelectorNode)
- [SequenceNode](https://mod-developers.nexon.com/apiReference/Misc/SequenceNode)
- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)
- [행동 트리를 활용한 AI 만들기](/docs?postId=562)

Update 2025-08-27 PM 04:56


# AIWanderComponent

주변을 배회하는 AI입니다. StateComponent가 없다면 자동으로 추가됩니다.

# Properties

##### inherited from AIComponent:

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Legacy 시스템 지원 여부를 설정합니다. 이전 시스템은 노드의 Running 상태를 구현하기 위해 ExclusiveExecutionWhenRunning 프로퍼티를 사용해야 합니다. Legacy 시스템은 더 이상 지원하지 않으며, 추후 삭제 예정입니다. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 값이 true면 행동 트리가 실행될 때 실행 관련 정보가 로그로 출력됩니다. 메이커 환경에서만 동작합니다. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 업데이트 권한입니다. Server와 Client 중 설정한 곳에서 실행됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from AIComponent:

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Action 노드를 생성합니다. 노드가 실행될 때 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| BTNodeType을 기반으로 Action 노드를 생성합니다. nodeType은 BTNodeType의 타입명입니다. onBehaveFunction가 nil이 아니라면 노드가 실행될 때 BTNodeType의 OnInit(), OnBehave() 함수 대신 onBehaveFunction으로 전달된 함수가 호출됩니다. onBehaveFunction의 매개 변수는 프레임당 시간인 delta입니다. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| node를 최상위 노드로 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# SeeAlso

- [StateComponent](https://mod-developers.nexon.com/apiReference/Components/StateComponent)

Update 2025-08-27 PM 04:56


# AreaParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

생성 영역을 조절할 수 있는 파티클 효과를 만드는 기능을 제공합니다. 눈, 비, 안개 같은 파티클이 해당합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티를 기준으로 생성 범위의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 생성 범위의 너비와 높이를 지정합니다. |

| [AreaParticleType](https://mod-developers.nexon.com/apiReference/Enums/AreaParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 생성할 파티클 타입을 설정합니다. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 반복 재생 여부를 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| 파티클을 재생합니다. |

| void Stop() |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

파티클을 재생, 정지시키고 `AreaSize`를 조절하는 예제입니다.

```
Property:
[Sync]
AreaParticleComponent ParticleComponent = nil

Method:
[client only]
void OnBeginPlay()
{
	self.ParticleComponent = self.Entity.AreaParticleComponent	 
	self.ParticleComponent.AreaSize.x = _UtilLogic:RandomDouble() * 2
	self.ParticleComponent.ParticleCount = 2
}

Event Handler:
[service: InputService]
HandleKeyDownEvent(KeyDownEvent event)
{
	----------------- Native Emitter Info ------------------
	-- Emitter: InputService
	-- Space: Client
	--------------------------------------------------------
	 
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.Q then
		self.ParticleComponent:Stop()
	elseif key == KeyboardKey.E then
		self.ParticleComponent:Play()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)
- [파티클 사용하기](/docs?postId=1036)
- [파티클 활용하기](/docs?postId=764)

Update 2025-12-02 PM 01:55


# AttackComponent

HitComponent와 연동해 공격 기능을 구현할 수 있는 인터페이스를 제공합니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) Attack([Shape](https://mod-developers.nexon.com/apiReference/Misc/Shape) shape, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| shape 영역 안에 있는 HitComponent의 OnHit(Entity, Integer, boolean, string, int32) 함수를 호출하고 HitEvent를 발생시킵니다. 공격 대상으로 판정된 HitComponent를 모두 반환합니다.<br>attackInfo는 사용자 정의 데이터로 공격을 직접 구현할 때 크리에이터가 원하는 용도에 맞게 활용할 수있습니다. 활용 시 함수를 재정의해야 합니다. |

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) Attack([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) size, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) offset, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| 사각형 영역을 지정할 수 있는 Attack 함수입니다. size는 사각형의 크기, offset은 엔티티 사각형의 중심 위치입니다. |

| void AttackFast([Shape](https://mod-developers.nexon.com/apiReference/Misc/Shape) shape, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| 반환 값이 없는 Attack 함수입니다. 불필요한 table 객체 생성을 줄여 월드 성능을 개선할 수 있습니다. |

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) AttackFrom([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) size, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| 사각형 영역을 지정할 수 있는 Attack 함수입니다. size는 사각형의 크기, position은 월드 좌표 기준 사각형의 중심 위치입니다. |

| boolean CalcCritical([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 크리티컬 공격 여부를 결정합니다. 기본 동작은 항상 false를 반환하므로 크리티컬 공격을 하지 않습니다. |

| integer CalcDamage([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 대미지 값을 결정합니다. 기본 동작은 항상 1을 반환합니다. |

| float GetCriticalDamageRate() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 크리티컬 공격일 경우, 기본 대미지 대비 몇 배의 대미지를 줄 것인지 결정합니다. 기본 동작은 항상 2를 반환합니다. |

| int32 GetDisplayHitCount([string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 한 번의 공격을 몇 개의 히트로 분할하여 표시할 지 결정합니다. 기본 동작은 항상 1을 반환합니다. |

| boolean IsAttackTarget([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| defender가 공격 대상인지 판단합니다. false를 반환하면 Attack(), AttackFrom(), AttackFast() 함수의 공격 대상에서 제외됩니다. defender의 StateComponent가 'DEAD'인 경우, 본인과 상대가 모두 플레이어면서 상대 PlayerComponent의 PVPMode 프로퍼티 값이 false인 경우에 false를 반환하는 것이 기본 동작입니다. |

| void OnAttack([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 이 엔티티가 공격할 때 호출되는 함수입니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [AttackEvent](https://mod-developers.nexon.com/apiReference/Events/AttackEvent) |
| --- |
| 엔티티가 공격할 때 발생하는 이벤트입니다. AttackComponent에서 발생합니다. |

# Examples

AttackComponent를 재정의 해 크리에이터가 공격 방식을 직접 구현할 수 있습니다.사각형 영역에 들어온 몬스터에게 공격을 가하는 예제입니다.

```
Method:
[server only]
void AttackNormal ()
{
	local attackSize = Vector2(1,1)
	local playerController = self.Entity.PlayerControllerComponent
	 
	if playerController ~= nil then
	l	ocal attackOffset = Vector2(0.5 * playerController.LookDirectionX, 0.5)
		self:Attack(attackSize, attackOffset, nil, CollisionGroups.Monster)
	end
}
 
-- 대미지 계산 방법을 정의합니다
override int CalcDamage (Entity attacker,Entity defender,string attackInfo)
{
	return 50
}
 
-- 크리티컬 대미지 계산 방법을 정의합니다.
override boolean CalcCritical (Entity attacker,Entity defender,string attackInfo)
{
	return _UtilLogic:RandomDouble() < 0.3
}
 
-- 크리티컬 데미지 발생 확률을 정의합니다.
override number GetCriticalDamageRate ()
{
	return 2
}

Event Handler:
[self]
HandlePlayerActionEvent (PlayerActionEvent event)
{
	-- Parameters
	local ActionName = event.ActionName
	local PlayerEntity = event.PlayerEntity
	--------------------------------------------------------
	if self:IsClient() then return end
	 
	if ActionName == "Attack" then
		self:AttackNormal()
	end
}
```

# SeeAlso

- [HitComponent](https://mod-developers.nexon.com/apiReference/Components/HitComponent)
- [PlayerControllerComponent](https://mod-developers.nexon.com/apiReference/Components/PlayerControllerComponent)
- [CollisionService](https://mod-developers.nexon.com/apiReference/Services/CollisionService)
- [공격과 피격](/docs?postId=206)

Update 2025-10-28 PM 02:21


# AvatarBodyActionSelectorComponent

아바타의 몸통 Action을 선택해 적용시키는 컴포넌트입니다.

# Properties

| [MapleAvatarBodyActionState](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarBodyActionState) ActionState |
| --- |
| 현재 Action 상태입니다. |

| [ReadOnlyDictionary<MapleAvatarBodyActionState, ReadOnlyList<MapleAvatarActionStateElement>>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyDictionary-2) StateActionDic ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 상태에 대한 정보값입니다. 추후 개선될 예정입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ActionStateChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ActionStateChangedEvent) |
| --- |
| 액션 상태가 변경될 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# AvatarFaceActionSelectorComponent

아바타의 얼굴 Action을 선택해 적용시키는 컴포넌트입니다.

# Properties

| [MapleAvatarFaceActionState](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarFaceActionState) ActionState |
| --- |
| 현재 Action 상태입니다. |

| float BlinkInterval |
| --- |
| 눈 깜빡임의 시간 간격으로, 초 단위를 사용합니다. |

| [ReadOnlyDictionary<MapleAvatarFaceActionState, ReadOnlyList<MapleAvatarActionStateElement>>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyDictionary-2) StateActionDic ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 상태에 대한 정보값입니다. 추후 개선될 예정입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# AvatarGUIRendererComponent

아바타 형태의 엔티티를 UI에 렌더링합니다.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 아바타 전체 색을 지정한 색상 값으로 변경합니다. |

| boolean FlipX |
| --- |
| X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY |
| --- |
| Y축을 기준으로 반전 여부를 결정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| float PlayRate |
| --- |
| 아바타 애니메이션 재생 속도를 지정할 수 있습니다. 0 이상의 값을 지원하며 숫자가 클수록 속도가 빨라집니다. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveAvatar |
| --- |
| 이미지의 비율, 크기, 피봇 보존 방식을 정의합니다. |

| boolean RaycastTarget |
| --- |
| true로 설정할 경우 화면 터치 또는 마우스 클릭 대상이 됩니다. 아바타 뒤에 가려진 UI는 화면 터치와 마우스 클릭 입력을 받지 못합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetAvatarRootEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타를 구성하는 최상위 엔티티인 AvatarRoot 엔티티를 가져옵니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetBodyEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타를 구성하는 Body 엔티티를 가져옵니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetFaceEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타를 구성하는 Face 엔티티를 가져옵니다. |

| void PlayEmotion([EmotionalType](https://mod-developers.nexon.com/apiReference/Enums/EmotionalType) emotionalType, float duration) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 감정 표현을 재생합니다. |

| void SetAvatarPartColor([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, float r, float g, float b, float a) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타의 category에 해당하는 파츠의 Color 값을 설정합니다. RGB 값을 사용합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

#### PlayEmotion

F1부터 F5 키를 눌러 감정 표현을 변경합니다.

```
Method:
[client only]
void ChangeEmotion (number emotionNumber)
{
	if emotionNumber == 1 then
	    self.Entity.AvatarGUIRendererComponent:PlayEmotion(EmotionalType.Vomit, 5)
	elseif emotionNumber == 2 then
	    self.Entity.AvatarGUIRendererComponent:PlayEmotion(EmotionalType.Glitter, 5)
	elseif emotionNumber == 3 then
	    self.Entity.AvatarGUIRendererComponent:PlayEmotion(EmotionalType.Hum, 5)
	elseif emotionNumber == 4 then
	    self.Entity.AvatarGUIRendererComponent:PlayEmotion(EmotionalType.Love, 5)
	elseif emotionNumber == 5 then
	    self.Entity.AvatarGUIRendererComponent:PlayEmotion(EmotionalType.Shine, 5)
	end
}

Event Handler:
[client only] [service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.F1 then
		self:ChangeEmotion(1)
	elseif key == KeyboardKey.F2 then
		self:ChangeEmotion(2)
	elseif key == KeyboardKey.F3 then
		self:ChangeEmotion(3)
	elseif key == KeyboardKey.F4 then
		self:ChangeEmotion(4)
	elseif key == KeyboardKey.F5 then
		self:ChangeEmotion(5)
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [아바타를 UI에서 표현하기](/docs?postId=953)

Update 2025-11-24 AM 11:42


# AvatarRendererComponent

아바타 형태의 엔티티를 렌더링합니다.

# Properties

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼의 Id를 지정합니다. |

| int32 OrderInLayer |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 아바타 애니메이션의 재생 속도입니다. |

| boolean ShowDefaultWeaponEffects ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 무기의 기본 이펙트 애니메이션과 기본 사운드 재생 여부를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Avatar의 SortingLayer를 설정합니다. 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetAvatarRootEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타를 구성하는 최상위 엔티티를 가져옵니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetBodyEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타의 몸통에 해당하는 엔티티를 가져옵니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetFaceEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 아바타의 얼굴에 해당하는 엔티티를 가져옵니다. |

| void PlayEmotion([EmotionalType](https://mod-developers.nexon.com/apiReference/Enums/EmotionalType) emotionalType, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 감정 표현을 재생합니다. |

| void SetAlpha(float alpha, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 아바타의 Alpha 값을 설정합니다. |

| void SetAvatarPartColor([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, float r, float g, float b, float a, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 아바타의 category에 해당하는 파츠의 Color 값을 설정합니다. RGB 값을 사용합니다. |

| void SetColor(float r, float g, float b, float a, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 아바타의 Color 값을 설정합니다. RGB 값을 사용합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# AvatarStateAnimationComponent

아바타의 상태 변화에 따라 재생될 애니메이션을 지정합니다.

# Properties

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Legacy 시스템을 지원할지를 설정합니다. 이전 시스템은 더 이상 지원하지 않으며, 추후 삭제 예정입니다. |

| [SyncDictionary<string, AvatarBodyActionElement>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) StateToAvatarBodyActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 상태의 이름과 AvatarBodyActionState가 매핑된 table입니다. IsLegacy의 값이 false일 때 사용됩니다. |

##### inherited from StateAnimationComponent:

| [SyncDictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 애니메이션의 이름과 AnimationClip이 매핑된 table입니다. IsLegacy의 값이 true일 때 사용됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| StateChangeEvent를 받았을 때 처리하는 함수입니다. 기본적으로 State에 매핑된 AnimationClip을 재생하는 AnimationClipEvent를 발생시킵니다. |

##### inherited from StateAnimationComponent:

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| StateChangeEvent를 받았을 때 처리하는 함수입니다. 기본적으로 State에 매핑된 AnimationClip을 재생하는 AnimationClipEvent를 발생시킵니다. |

| void RemoveActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key) |
| --- |
| StateToAvatarBodyActionSheet에서 key에 해당하는 요소를 제거합니다. IsLegacy 값이 true면 ActionSheet에서 요소를 제거합니다. |

| void SetActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) animationClipRuid) |
| --- |
| StateToAvatarBodyActionSheet에 요소를 추가합니다. 요소로 추가되는 AvatarBodyActionElement 객체의 AvatarBodyActionStateName 프로퍼티 값은 animationClipRuid, PlayerRate 프로퍼티 값은 1이 됩니다. IsLegacy의 값이 true면 ActionSheet에 요소를 추가합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) StateStringToAnimationKey([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| State에 매핑된 Animation의 이름을 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [BodyActionStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/BodyActionStateChangeEvent) |
| --- |
| BodyAction의 상태가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# BackgroundComponent

맵의 배경을 변경하고 관리합니다. 단일 색상 배경, 메이플스토리 배경, 웹 이미지 배경, 그라데이션 배경을 사용할 수 있습니다.

# Properties

| [SyncDictionary<string, BackgroundPieceDataElement>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) BackgroundPieces ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 배경의 Background Piece 목록입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 최하단 색상 값입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) MiddleBottomColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 중하단 색상 값입니다. |

| float MiddleBottomRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 중하단 색상의 기준 위치입니다. 0.01 이상, MiddleTopRatio - 0.01 이하의 값을 설정할 수 있습니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) MiddleTopColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 중상단 색상 값입니다. |

| float MiddleTopRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 중상단 색상의 기준 위치입니다. MiddleBottomRatio + 0.01 이상, 0.99 이하의 값을 설정할 수 있습니다. |

| float ScrollRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Type이 Template일 때 배경의 움직이는 속도를 조절할 수 있습니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) SolidColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Type이 SolidColor일 때의 색상 값입니다. 배경을 변경하려면 ChangeBackgroundBySolidColor(Color) 함수를 사용하세요. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) TemplateRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Type이 Template일 때의 RUID입니다. 배경을 변경하려면 ChangeBackgroundByTemplateRUID(string) 함수를 사용하세요. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 그라데이션 배경의 최상단 색상 값입니다. |

| [BackgroundType](https://mod-developers.nexon.com/apiReference/Enums/BackgroundType) Type ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 배경의 타입입니다. ChangeBackgroundBy류 함수를 사용해 배경을 바꿀 수 있습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) WebUrl ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Type이 Web일 때의 웹 이미지 배경의 주소입니다. 배경을 변경하려면 ChangeBackgroundByWebUrl(string) 함수를 사용하세요. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeBackgroundByGradient([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) top, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) middleTop, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) middleBottom, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) bottom, float middleTopRatio, float middleBottomRatio) |
| --- |
| 배경을 그라데이션 색상들로 변경합니다. |

| void ChangeBackgroundBySolidColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) value) |
| --- |
| 배경을 단색으로 변경합니다. |

| void ChangeBackgroundByTemplateRUID([string](https://mod-developers.nexon.com/apiReference/Lua/string) value) |
| --- |
| 배경을 지정한 템플릿 배경 RUID로 변경합니다. |

| void ChangeBackgroundByWebUrl([string](https://mod-developers.nexon.com/apiReference/Lua/string) value) |
| --- |
| 배경을 웹 이미지로 변경합니다. |

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| void SetBackgroundPieceColor([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| pieceName에 해당하는 Background Piece를 찾아 색상을 설정합니다. |

| void SetBackgroundPieceEnable([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, boolean enable) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| pieceName에 해당하는 Background Piece를 찾아 활성화 여부를 설정합니다. |

| void SetBackgroundPiecePosition([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| pieceName에 해당하는 Background Piece를 찾아 위치를 설정합니다. |

| void SetBackgroundPieceRatio([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ratio) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| pieceName에 해당하는 Background Piece를 찾아 화면 이동 비율을 설정합니다. 화면 이동 비율은 카메라가 이동할 때 상대적으로 얼마큼 이동할 지 나타내는 값입니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

다양한 방법으로 배경을 변경하는 예제입니다.

```
Method:
[server only]
void ChangeBackgroundExample ( )
{
	local background = self.Entity.BackgroundComponent
	background:ChangeBackgroundByTemplateRUID("794ad8421e2543d8a6d2c70307637450")
	background:ChangeBackgroundBySolidColor(Color.white)
	background:ChangeBackgroundByWebUrl("Http://WebUrl")
}
```

# SeeAlso

- [배경 설정하기](/docs?postId=768)

Update 2025-12-02 PM 01:55


# BaseParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)![custom](https://img.shields.io/static/v1?label=&amp;message=Abstract&amp;color=darkkhaki)

파티클 효과를 만드는 ParticleComponent들의 부모 컴포넌트입니다.

# Properties

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 반복 재생 여부를 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void Play() |
| --- |
| 파티클을 재생합니다. |

| void Stop() |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# SeeAlso

- [AreaParticleComponent](https://mod-developers.nexon.com/apiReference/Components/AreaParticleComponent)
- [BasicParticleComponent](https://mod-developers.nexon.com/apiReference/Components/BasicParticleComponent)
- [SpriteParticleComponent](https://mod-developers.nexon.com/apiReference/Components/SpriteParticleComponent)
- [AreaParticleType](https://mod-developers.nexon.com/apiReference/Enums/AreaParticleType)
- [BasicParticleType](https://mod-developers.nexon.com/apiReference/Enums/BasicParticleType)
- [SpriteParticleType](https://mod-developers.nexon.com/apiReference/Enums/SpriteParticleType)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)

Update 2025-08-27 PM 04:56


# BasicParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

기본 파티클의 설정 및 제어 기능을 제공합니다.

# Properties

| [BasicParticleType](https://mod-developers.nexon.com/apiReference/Enums/BasicParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 생성할 파티클의 타입을 설정합니다. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 반복 재생 여부를 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| 파티클을 재생합니다. |

| void Stop() |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

파티클을 재생, 정지시키고 파티클 프로퍼티를 제어하는 예제입니다.

```
Property:
[Sync]
BasicParticleComponent ParticleComponent = nil

Method:
[client only]
void OnBeginPlay()
{
	self.ParticleComponent = self.Entity.BasicParticleComponent
	 
	self.ParticleComponent.PlaySpeed = _UtilLogic:RandomDouble() * 2
	self.ParticleComponent.ParticleCount = 2
}

Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	----------------- Native Emitter Info ------------------
	-- Emitter: InputService
	-- Space: Client
	--------------------------------------------------------
	 
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.Q then
		self.ParticleComponent:Stop()
	elseif key == KeyboardKey.E then
		self.ParticleComponent:Play()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)
- [파티클 사용하기](/docs?postId=1036)
- [파티클 활용하기](/docs?postId=764)

Update 2025-12-02 PM 01:55


# ButtonComponent

UI 버튼 기능을 제공합니다.

# Properties

| TransitionColorSet Colors ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 버튼의 상태별 색상을 정의합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| TransitionRUIDSet ImageRUIDs ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 버튼의 상태별 이미지를 정의합니다. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되었는지 여부를 나타냅니다. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) KeyCode |
| --- |
| 버튼을 누르면 지정한 keyCode를 누른 것처럼 동작합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [TransitionType](https://mod-developers.nexon.com/apiReference/Enums/TransitionType) Transition |
| --- |
| 버튼의 상태 전환 효과를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ButtonClickEditorEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonClickEditorEvent) |
| --- |
| Button 클릭 시 발생하는 에디터 이벤트입니다. |

| [ButtonClickEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonClickEvent) |
| --- |
| Button 클릭 시 발생하는 이벤트입니다. |

| [ButtonPressedEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonPressedEvent) |
| --- |
| Button의 상태가 Pressed가 됐을 때 발생하는 이벤트입니다. |

| [ButtonStateChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonStateChangeEditorEvent) |
| --- |
| Button의 상태가 변경될 때 발생하는 에디터 이벤트입니다. |

| [ButtonStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonStateChangeEvent) |
| --- |
| Button의 상태가 변경될 때 발생하는 이벤트입니다. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

버튼을 누르고 있을 때 버튼 이미지가 변경되고 이미지의 색상이 변경되는 예제입니다. `ButtonComponent`가 있는 Entity에 이 Component를 추가해 테스트할 수 있습니다.

```
Property:
[None]
boolean IsButtonDown = false
[None]
number RedVal = 0
[None]
number TimerID = 0
[None]
string MonsterRUID = "000000"
[None]
string OriginalRUID = ""
  
  
Method:
[client only]
void OnBeginPlay ()
{
	self.OriginalRUID = self.Entity.SpriteGUIRendererComponent.ImageRUID
}
  
void CancelHoldButton ()
{
	_TimerService:ClearTimer(self.TimerID)
	self.Entity.SpriteGUIRendererComponent.ImageRUID = self.OriginalRUID
	self.IsButtonDown = false
	self.RedVal = 0
}
  
Event Handler:
[self]
HandleButtonPressedEvent (ButtonPressedEvent event)
{
	-- Parameters
	-------------------------------------------------------
	if self.IsButtonDown then
		self:CancelHoldButton()
	else
		self.Entity.SpriteGUIRendererComponent.ImageRUID = self.MonsterRUID
	end
	
	self.IsButtonDown = true
	
	local AddMoreRed  = function()
		if self.RedVal <= 1 then
			self.RedVal = self.RedVal + 0.2
			self.Entity.ButtonComponent.Colors.PressedColor = Color(self.RedVal, 0, 0)
		end
	end
	self.TimerID = _TimerService:SetTimerRepeat(AddMoreRed, 0.3)
}
  
[self]
HandleButtonClickEvent (ButtonClickEvent event)
{
	-- Parameters
	local Entity = event.Entity
	--------------------------------------------------------
	self.Entity.SpriteGUIRendererComponent.ImageRUID = self.OriginalRUID
	self.IsButtonDown = false

}
  
[self]
HandleButtonStateChangeEvent (ButtonStateChangeEvent event)
{
	-- Parameters
	local state = event.state
	-------------------------------------------------------
	if state == ButtonState.Released then
		self:CancelHoldButton()
	end
}
```

# SeeAlso

- [SpriteGUIRendererComponent](https://mod-developers.nexon.com/apiReference/Components/SpriteGUIRendererComponent)
- [ButtonState](https://mod-developers.nexon.com/apiReference/Enums/ButtonState)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [기본 UI 컴포넌트](/docs?postId=744)

Update 2025-12-02 PM 01:55


# CameraComponent

Entity를 바라보는 카메라 기능을 제공합니다. 카메라 간 화면 전환은 CameraService를 이용하세요.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CameraOffset |
| --- |
| 카메라의 위치를 설정합니다. 월드 좌표 기준입니다. |

| boolean ConfineCameraArea |
| --- |
| 카메라가 비추는 범위를 맵의 발판 영역으로만 제한합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Damping |
| --- |
| 카메라가 대상을 추적하는 동안 대상이 SoftZone에 들어갈 때 카메라가 반응하는 속도를 조정합니다. 값이 작을수록 더욱 더 빠르게 반응합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DeadZone |
| --- |
| DeadZone 영역을 설정합니다. 카메라가 타겟을 유지하는 프레임 영역입니다. |

| float DutchAngle |
| --- |
| 카메라의 회전 값을 설정합니다. |

| boolean IsAllowZoomInOut |
| --- |
| 줌 기능 사용 여부를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LeftBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 카메라 제한 영역의 좌하단 값입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RightTop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 카메라 제한 영역의 우상단 값입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ScreenOffset |
| --- |
| 대상을 기준으로 한 전체 스크린의 비율 값입니다. 0부터 1 사이의 값을 사용할 수 있고, 값이 0.5일 경우 카메라가 중앙에 위치합니다. ConfineCameraArea가 false인 경우 사용 가능합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SoftZone |
| --- |
| SoftZone 영역을 설정합니다. 대상이 프레임의 영역에 들어오면 카메라가 방향을 바꾸어 DeadZone으로 되돌립니다. |

| boolean UseCustomBound ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 카메라 제한 영역을 직접 정의해 사용할지 여부를 나타냅니다. LeftBottom, RightTop 프로퍼티를 사용해 정의할 수 있습니다. false일 시 속해있는 맵의 맵 영역을 사용합니다. 맵 영역이 기본이라면 그 영역을 바탕으로 보정된 영역을 카메라 제한 영역으로 사용합니다. |

| float ZoomRatio |
| --- |
| 줌 비율을 설정합니다. 값은 ZoomRatioMin 이상, ZoomRationmax 이하로만 설정할 수 있습니다. 단위는 백분율입니다. |

| float ZoomRatioMax |
| --- |
| 카메라의 줌 비율의 최댓값을 설정합니다. 500보다 큰 값으로 설정할 수 없습니다. 단위는 백분율입니다. |

| float ZoomRatioMin |
| --- |
| 카메라의 줌 비율의 최소값을 설정합니다. 30보다 작은 값으로 설정할 수 없습니다. 단위는 백분율입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| Vector2, Vector2 GetBound() |
| --- |
| LeftBottom, RightTop으로 구성된 카메라 제한 영역을 가져옵니다. |

| void SetZoomTo(float percent, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 주어진 시간(초)동안 카메라를 확대합니다. |

| void ShakeCamera(float intensity, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 주어진 시간 동안 카메라를 진동시킵니다. 초 단위를 사용합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

4초 뒤에 줌 배율을 300%로 설정하는 코드입니다.

```
Method:
[server only]
void OnBeginPlay ( )
{
	local zoom = function()
	    self.Entity.CameraComponent:SetZoomTo(300, 2)
	end
	_TimerService:SetTimerOnce(zoom, 4)
}
```

# SeeAlso

- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [CameraService를 활용한 카메라 제어](/docs?postId=118)

Update 2025-08-27 PM 04:56


# CanvasGroupComponent

자식 UI Entity의 공통 속성을 한 곳에서 제어할 수 있는 기능을 제공합니다.

# Properties

| boolean BlocksRaycasts |
| --- |
| true로 설정할 경우 자식 UI는 화면 터치 또는 마우스 클릭 입력을 받을 수 있습니다. 뒤에 가려진 UI는 입력을 받지 못합니다. |

| float GroupAlpha |
| --- |
| 자식 UI의 투명도를 설정합니다. 0 이상, 1 이하의 값을 설정할 수 있습니다. |

| boolean Interactable |
| --- |
| 자식 UI의 상호작용 가능 여부를 설정합니다. false로 설정할 경우 사용자 입력을 받아도 아무런 동작을 하지 않습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# SeeAlso

- [기본 UI 컴포넌트](/docs?postId=744)
- [UI 제작하기](/docs?postId=64)

Update 2025-10-28 PM 02:21


# ChatBalloonComponent

말풍선을 표시해주고 관련 기능을 설정할 수 있는 기능을 제공합니다.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Message 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean ArrowChatEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선 꼬리 이미지를 제공 여부를 설정합니다. |

| boolean AutoShowEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선이 자동으로 나타나게 할지 여부를 설정합니다. 값이 true일 경우 ShowDuration 동안 말풍선이 보이고 HideDuration 동안 말풍선이 사라지는 동작을 반복합니다. 플레이어의 ChatModeEnabled 값이 true일 경우 자동 말풍선 기능이 동작하지 않습니다. |

| float BalloonScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선의 크기를 설정합니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ChatBalloonRUID |
| --- |
| 말풍선의 종류를 설정합니다. |

| boolean ChatModeEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 채팅 연동 여부를 설정합니다. 플레이어만 동작합니다. 값이 true면 채팅 내용이 말풍선으로 출력되며, AutoShowEnabled가 true여도 자동 말풍선 기능은 동작하지 않습니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선의 글자색을 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) FontOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선 내 글자 위치를 상하좌우로 조정합니다. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선 내의 글자 크기를 설정합니다. |

| float HideDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| AutoShowEnabled의 값이 true일 때 말풍선이 보이지 않는 시간을 설정합니다. 설정한 시간이 지난 후에 말풍선이 다시 나타납니다. 초 단위를 사용합니다. |

| boolean IsRichText ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 리치 텍스트 사용 여부를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Message ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선에 표시할 내용을 설정합니다. nil 또는 빈 문자열일 경우 말풍선이 나타나지 않습니다. |

| float Offset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선의 위치 offset입니다. 말풍선의 위치를 위, 아래로 조정할 수 있습니다. |

| float ShowDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선이 떠 있는 시간을 설정합니다. 설정한 시간이 지나면 말풍선이 자동으로 사라집니다. 초 단위를 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ChatBalloonEvent](https://mod-developers.nexon.com/apiReference/Events/ChatBalloonEvent) |
| --- |
| 말풍선이 나타날 때 생기는 이벤트입니다. |

# Examples

다음은 일정 간격마다 대사를 말하는 NPC를 만드는 예제입니다. ChatBalloon이 출력되는 간격에 맞추어 대사를 교체해 줍니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	local chatBalloon = self.Entity.ChatBalloonComponent
	 
	local npcTalks = {}
	table.insert(npcTalks, "NPC Talk 1")
	table.insert(npcTalks, "NPC Talk 2")
	table.insert(npcTalks, "NPC Talk 3")
	table.insert(npcTalks, "NPC Talk 4")
	 
	local setMessage = function()
		local idx = _UtilLogic:RandomIntegerRange(1, #npcTalks)
		chatBalloon.Message = npcTalks[idx]
	end
	 
	local interval = chatBalloon.ShowDuration + chatBalloon.HideDuration
	_TimerService:SetTimerRepeat(setMessage, interval)
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [말풍선 만들기](/docs?postId=119)

Update 2025-10-28 PM 02:21


# ChatComponent

플레이어끼리 소통을 할 수 있는 채팅 기능을 지원해주는 컴포넌트입니다.

# Properties

| float ChatEmotionDuration |
| --- |
| 아바타 감정 표현 지속 시간을 설정합니다. |

| boolean EnableVoiceChat ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 보이스 채팅 버튼의 표시 및 사용 여부를 설정합니다. |

| boolean Expand |
| --- |
| 채팅창을 펼칠 수 있는 기능입니다. |

| boolean HideWorldChatButton ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 월드 채팅 버튼을 숨기는 기능입니다. |

| boolean MessageAlignBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 월드 채팅 메세지를 하단을 기준으로 정렬하는 기능입니다. |

| boolean UseChatBalloon |
| --- |
| 플레이어의 채팅 메세지를 말풍선으로 표현할 수 있는 기능입니다. |

| boolean UseChatEmotion |
| --- |
| 채팅으로 아바타 감정 표현을 사용할 수 있는 기능입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ChatEvent](https://mod-developers.nexon.com/apiReference/Events/ChatEvent) |
| --- |
| 대화가 입력되었을 때 발생하는 이벤트입니다. |

# Examples

`EmotionalType`의 텍스트 길이만큼 `ChatEmotionDuration`과 `ChatBalloonComponent`의 `ShowDuration`을 변경하는 예제입니다. `ChatEvent`를 통해 EmotionalType이 메시지에 포함되었는지 확인합니다.

```
Event Handler:
[self]
HandleChatEvent (ChatEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: ChatComponent
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local Message = event.Message
	local SenderName = event.SenderName
	local UserId = event.UserId
	---------------------------------------------------------
	 
	local localId = _UserService.LocalPlayer.OwnerId
	if string.compare(localId, UserId) == false then
		return
	end
	 
	local lowerMessage = string.lower(Message)
	 
	local findEmotion = EmotionalType.Invalid
	local len = 23
	for i = 1, len do
		local key = string.lower(tostring(EmotionalType.CastFrom(i)))
	 
		if lowerMessage:find(key, 1, true ) then
			log("onChat Find : ", key)
			findEmotion = EmotionalType.CastFrom(i)
		end
	end
	 
	if findEmotion == EmotionalType.Invalid then
		return
	end
	 
	local chatComponent = self.Entity.ChatComponent
	local duration = #tostring(findEmotion)
	if chatComponent then
		chatComponent.UseChatEmotion = true
		chatComponent.ChatEmotionDuration = duration
	end
	 
	local balloonComponent = _UserService.LocalPlayer.ChatBalloonComponent
	if balloonComponent then
		balloonComponent.ShowDuration = duration
	end
	 
	 
	log("OnChatEvent ", Message, " duration :", duration)
}
```

# SeeAlso

- [ChatBalloonComponent](https://mod-developers.nexon.com/apiReference/Components/ChatBalloonComponent)
- [EmotionalType](https://mod-developers.nexon.com/apiReference/Enums/EmotionalType)
- [ChatEvent](https://mod-developers.nexon.com/apiReference/Events/ChatEvent)
- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)

Update 2025-08-27 PM 04:56


# ClimbableComponent

등반 액션을 할 수 있는 영역을 지정합니다.

# Properties

| boolean AllowHorizontalMove |
| --- |
| 자유로운 등반 이동 여부를 설정합니다. true일 경우 X, Y 축 모두 이동 가능합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티를 기준으로 충돌체 직사각형의 중심점 위치를 설정합니다. IsUseDefaultObjectBoxSize를 false로 설정해 BoxOffset을 적용할 수 있습니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌체 직사각형의 너비와 높이를 지정합니다. IsUseDefaultObjectBoxSize를 false로 설정해 BoxSize를 적용할 수 있습니다. |

| [ClimbableType](https://mod-developers.nexon.com/apiReference/Enums/ClimbableType) ClimbableAnimationType |
| --- |
| 등반 시 애니메이션 타입을 결정합니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 등반 영역의 충돌 그룹입니다. |

| boolean IsUseDefaultObjectBoxSize |
| --- |
| 등반 영역 설정을 스프라이트 크기에 맞출지 설정합니다. 메이커 편집 모드에서만 동작합니다. true일 경우에 Boxsize, BoxOffset을 자동으로 스프라이트 크기에 맞게 설정합니다. false일 경우 Boxsize, BoxOffset을 임의로 설정할 수 있습니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SpeedFactor |
| --- |
| 등반 시 X, Y축 속력에 곱해지는 가중치입니다. 값이 클수록 이동 속력이 빨라집니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

`TriggerEnterEvent` 발생 시 충돌한 엔티티의 이름에 따라 `ClimbableComponent`의 `BoxSize`와 `SpeedFactor`가 변경되는 예제입니다.

```
Method:
[server]
void ChangeBoxSize (Vector2 boxSize)
{
	local entity = _EntityService:GetEntityByPath(EntityPath)
	entity.ClimbableComponent.BoxSize = boxSize
}

[client]
void ChangeSpeedVector (Vector2 speedVector)
{
	local entity = _EntityService:GetEntityByPath(EntityPath)
	entity.ClimbableComponent.SpeedFactor = speedVector
}

Event Handler:
[server only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TriggerComponent
	-- Space: Server, Client
	---------------------------------------------------------
	 
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	---------------------------------------------------------
	 
	if TriggerBodyEntity.Name == "Name" then
		self:ChangeBoxSize(Vector2(3, 10))
	elseif TriggerBodyEntity.Name == "Name" then
		self:ChangeSpeedVector(Vector2(1, 2))
	end
}
```

# SeeAlso

- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [사다리와 로프 활용하기](/docs?postId=809)

Update 2025-08-27 PM 04:56


# ClimbableSpriteRendererComponent

Climbable Sprite를 설정하고 월드 상에 이미지를 표시합니다.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ClipName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sprite에 색깔을 씌웁니다. 기본 흰색일 경우에 원본의 색이 나옵니다. |

| [SpriteDrawMode](https://mod-developers.nexon.com/apiReference/Enums/SpriteDrawMode) DrawMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 스프라이트가 그려지는 방식을 설정합니다. Tiled를 사용할 수 있습니다. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트의 X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트의 Y축을 기준으로 반전 여부를 결정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 랜더러가 사용할 머티리얼의 Id를 지정합니다. |

| boolean NeedGizmo ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 스프라이트의 Gizmo를 표시합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| RenderSettingType RenderSetting ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 사다리 몸통에 사용할 Sprite RUID를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUIDHead ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 사다리 머리에 사용할 Sprite RUID를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUIDTail ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 사다리 꼬리에 사용할 Sprite RUID를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TiledSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Y값을 변경해 사다리의 세로 높이를 설정 할 수 있습니다. X값 변경은 지원하지 않습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| void ResetColliderBox() |
| --- |
| 사다리의 스프라이트 크기에 맞게 충돌 영역을 초기화합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [EmbededSpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeFrameEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerChangeFrameEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeStateEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerChangeStateEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerEndEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerEndEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerStartEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerStartEvent를 사용하세요. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeFrameEvent) |
| --- |
| 스프라이트 애니메이션의 프레임이 바뀔 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeStateEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. |

| [SpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndEvent) |
| --- |
| 스프라이트 애니메이션 재생이 끝날 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndFrameEvent) |
| --- |
| 스프라이트 애니메이션이 마지막 프레임을 재생할 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartEvent) |
| --- |
| 스프라이트 애니메이션 재생이 시작될 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartFrameEvent) |
| --- |
| 스프라이트 애니메이션이 첫번째 프레임을 재생할 때 발생하는 이벤트입니다. |

# Examples

`TriggerEnterEvnet` 발생 시 충돌한 엔티티의 이름에 따라 `ClimbableRendererComponent`의 `TiledSize`를 변경하고 스프라이트 크기에 맞게 충돌 박스 크기를 변경하는 예제입니다.

```
Method:
[server]
void ResizeRenderer (Vector2 tiledSize)
{
	local entity = _EntityService:GetEntityByPath(EntityPath)
	entity.ClimbableSpriteRendererComponent.TiledSize = tiledSize
	entity.ClimbableSpriteRendererComponent:ResetColliderBox()
}

Event Handler:
[server only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TriggerComponent
	-- Space: Server, Client
	---------------------------------------------------------
	 
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	---------------------------------------------------------
	 
	if TriggerBodyEntity.Name == "Name" then
		self:ResizeRenderer(Vector2(1, 5))
	end
}
```

# SeeAlso

- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [사다리와 로프 활용하기](/docs?postId=809)

Update 2025-08-27 PM 04:56


# Component

모든 컴포넌트의 부모 컴포넌트입니다. 컴포넌트 기본 기능을 제공합니다.

# Properties

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# CostumeManagerComponent

플레이어가 장착한 의상, 무기 등의 정보를 관리합니다.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomBodyEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 피부 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCapeEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 망토 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCapEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 모자 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCoatEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 코트 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEarAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 귀 장식 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEarEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 귀 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEyeAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 눈 장식 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomFaceAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 얼굴 장식 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomFaceEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 얼굴 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomGloveEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 장갑 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomHairEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 헤어스타일 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomLongcoatEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 롱코트 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomOneHandedWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 한손무기 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomPantsEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 하의 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomShoesEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 신발 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomSubWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 보조무기 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomTwoHandedWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 두손무기 |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) DefaultEquipUserId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 대상 유저의 장비를 복제한 이후 커스텀 장비들이 적용됩니다. 접속 중이지 않은 유저도 지정할 수 있으며, 지정 후 대상 유저의 장비가 변경된 경우는 반영되지 않습니다. |

| [ReadOnlyList<MapleAvatarItemData>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyList-1) EquippedItems ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 착용 중인 장비들의 정보입니다. 스크립트에서 수정할 수 없습니다. |

| boolean UseCustomEquipOnly ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 유저가 장착한 기본 의상을 사용하지 않고, 스크립트를 이용해 지정한 의상만을 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetEquip([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 현재 장착 중인 itemRUID를 반환합니다. 프로퍼티를 이용한 접근이 가능합니다. |

| void SetEquip([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, [string](https://mod-developers.nexon.com/apiReference/Lua/string) itemRUID) |
| --- |
| 입력한 itemRUID에 해당하는 아이템을 장착합니다. 장비를 해제하려면 빈 문자열을 입력합니다. 프로퍼티를 이용한 변경이 가능합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [DestroyMapleCostumeEvent](https://mod-developers.nexon.com/apiReference/Events/DestroyMapleCostumeEvent) |
| --- |
| CostumeManagerComponent가 삭제될 때 발생하는 이벤트입니다. |

| [InitMapleCostumeEvent](https://mod-developers.nexon.com/apiReference/Events/InitMapleCostumeEvent) |
| --- |
| 장비 착용 상태에 변화가 있을 때 발생하는 이벤트입니다. |

# Examples

다음 예제는 플레이어와 다른 Entity가 닿았을 때 플레이어의 헤어를 변경하고 RUID를 출력합니다.

```
Event Handler:
[self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	--Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	if self:IsServer() == true then
	    return
	end
	     
	local costumeManager = TriggerBodyEntity.CostumeManagerComponent
	costumeManager:SetEquip(MapleAvatarItemCategory.Hair, "000000")
	log (costumeManager:GetEquip(MapleAvatarItemCategory.Hair))
}
```

# SeeAlso

- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)

Update 2025-08-27 PM 04:56


# CustomFootholdComponent

사용자 정의 발판을 만드는 기능을 제공합니다. 발판의 모양과 속성을 설정할 수 있습니다

# Properties

| [SyncList<List<Vector2>>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) edgeLists ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 발판을 나타냅니다. 하나의 발판은 여러 개의 연결점으로 구성됩니다. |

| float FootholdDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| RigidbodyComponent가 포함된 엔티티가 발판 위에 있을 때 적용되는 마찰력입니다. 값이 클수록 빠르게 감속합니다. |

| float FootholdForce ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| RigidbodyComponent가 포함된 엔티티가 발판 위에 있을 때 가해지는 힘입니다. 값이 양수면 오른쪽, 음수면 왼쪽으로 움직입니다. |

| float FootholdWalkSpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| RigidbodyComponent가 포함된 엔티티가 발판 위에 있을 때 이동 속도에 곱해지는 계수입니다. 값이 클수록 이동 속도가 빨라집니다. |

| boolean IsBlockVerticalLine ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티가 해당 타일맵의 세로 발판에 막힐지 결정합니다. RigidbodyComponent를 가진 엔티티만 해당합니다. |

| boolean IsDynamicFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 값이 true인 경우 플레이 중 발판을 이동시키거나 모양을 변경할 수 있습니다. 발판의 위치, 모양을 자주 변경하면 월드 성능 저하의 요인이 될 수 있으므로 사용 시 주의가 필요합니다. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| true일 경우 Physics 기능을 사용하는 Dynamic 강체(PhysicRigidbody)와 충돌할 수 있습니다. |

| [RigidbodyMovementOptionType](https://mod-developers.nexon.com/apiReference/Enums/RigidbodyMovementOptionType) RigidbodyMovementOption ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이동하는 발판 위에서 RigidbodyComponent가 상대적으로 어떻게 움직일지 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

`edgeList`를 수정해 발판 모양을 변경할 수 있습니다. 발판 변경 시 아래 유의사항을 참고 해야합니다.

1. `IsDynamicFoothold`가 true인 경우에만 변경 사항이 적용됩니다.
2. 새 List를 만들어 덮어써야 변경 사항이 적용됩니다. 아래 예제 코드를 참고하세요.

발판 일부를 변경하는 예제입니다.

```
local target = _EntityService:GetEntityByPath(EntityPath)
 
-- 요소의 값을 수정하면 변경 사항이 발판에 적용되지 않습니다.
-- target.CustomFootholdComponent.edgeLists[1][1] = Vector2(3, 0)
 
-- 새 List를 만들어 덮어써야 합니다.
local list = {Vector2(3,0), Vector2(2,3)}
target.CustomFootholdComponent.edgeLists[1] = list
```

발판 전체를 대체하는 예제입니다.

```
local listlist = {{Vector2(0,2), Vector2(3,2)},{Vector2(4,5), Vector2(3,8)}}
local target = _EntityService:GetEntityByPath(EntityPath)
target.CustomFootholdComponent.edgeLists = listlist
```

# SeeAlso

- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [발판 만들기](/docs?postId=71)
- [이동 발판 만들기](/docs?postId=579)

Update 2025-12-02 PM 01:55


# DamageSkinComponent

대미지를 시각적으로 표현하는 대미지 스킨을 구성합니다. 대미지 스킨의 형식은 공격자 엔티티의 DamageSkinSettingComponent에서 지정합니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# DamageSkinSettingComponent

공격 시 재생할 대미지 스킨의 형식을 설정합니다. 피격 엔티티는 DamageSkinSpawnerComponent가 있어야 합니다.

# Properties

| float Alpha ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨의 Alpha 값을 설정합니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) DamageSkinId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨의 종류를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DamageSkinScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨의 크기를 설정합니다. |

| float DelayPerAttack ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 분할 재생 간격을 설정합니다. 초 단위를 사용합니다. AttackComponent의 GetDisplayHitCount(attackInfo) 함수를 참고하세요. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨의 재생 속도를 설정합니다. |

| [DamageSkinTweenType](https://mod-developers.nexon.com/apiReference/Enums/DamageSkinTweenType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨의 움직임 형태를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# DamageSkinSpawnerComponent

Entity에 피격 이벤트가 발생했을 때 대미지 스킨을 출력합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DamageSkinOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 대미지 스킨이 출력되는 위치를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

현재 HitComponent의 ColliderOffset과 BoxSize에 맞춰 상단 중앙에 대미지를 출력하는 예제입니다.

```
Method:
[server only]
void OnBeginPlay()
{
	local damageSkinSpawnerComponent = self.Entity.DamageSkinSpawnerComponent
	
	if damageSkinSpawnerComponent == nil then
	    return
	end
	
	local transformComponent = self.Entity.TransformComponent
	
	if transformComponent == nil then
		return
	end
	
	local hitComponent = self.Entity.HitComponent
	
	if hitComponent == nil then
	    return
	end 
	
	local weight = self.Entity.TransformComponent.Scale:ToVector2()
	 
	local hitBoxOffset = hitComponent.ColliderOffset * weight
	local hitBoxSize = hitComponent.BoxSize * weight
	
	local damageOffset = Vector2(hitBoxOffset.x, hitBoxOffset.y + hitBoxSize.y / 2)
	
	damageSkinSpawnerComponent.DamageSkinOffset = damageOffset
}
```

# SeeAlso

- [HitComponent](https://mod-developers.nexon.com/apiReference/Components/HitComponent)
- [TransformComponent](https://mod-developers.nexon.com/apiReference/Components/TransformComponent)
- [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2)

Update 2025-08-27 PM 04:56


# DirectionSynchronizerComponent

부모 엔티티의 방향을 따라갑니다. 부모 엔티티에서 ChangedLookAtEvent가 발생할 경우 현재 엔티티의 방향도 변경됩니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

아래 예제는 object-1모델을 워크스페이스에 추가한 뒤에 실행 가능합니다. DefaultPlayer에 아래와 같은 Component를 추가하면 플레이어의 방향과 같은 곳을 바라보는 오브젝트를 생성할 수 있습니다.

```
Method:
[server only]
void OnBeginPlay()
{
	local modelId = "Model Entry ID"
	local attachedEntity = _SpawnService:SpawnByModelId(modelId, "attachedEntity", Vector3.zero, self.Entity, "")
	attachedEntity:AddComponent("DirectionSynchronizerComponent")
}
```

# SeeAlso

- [SpawnService](https://mod-developers.nexon.com/apiReference/Services/SpawnService)

Update 2025-08-27 PM 04:56


# DistanceJointComponent

DistanceJoint를 생성, 삭제할 수 있습니다. 연결된 강체 간 거리를 일정하게 유지합니다.

# Properties

| [SyncList<DistanceJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Joint 정보를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, float length) |
| --- |
| DistanceJoint를 추가합니다. 성공 시 index, 실패 시 -1을 반환합니다 |

| void DestroyJoint(int32 index) |
| --- |
| 순번이 index에 해당하는 Joint를 제거합니다. |

| int32 GetJointsCount() |
| --- |
| joint 수를 반환합니다. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| 순번이 index에 해당하는 Joint의 CollideConnected 값을 설정합니다. |

| void SetLength(int32 index, float length) |
| --- |
| 순번이 index에 해당하는 Joint의 Length 값을 설정합니다. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorA 값을 설정합니다. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorB 값을 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

#### AddJoint

R을 눌러 Joint를 추가해 떨어지는 공을 잡는 예제입니다.

```
Function:
[server]
void Connect ()
{
	local ball = _EntityService:GetEntityByPath("/maps/map01/Ball")
	
	self.Entity.DistanceJointComponent:AddJoint(
		ball,
		Vector2.zero,
		Vector2.zero,
		Vector2.Distance(ball.TransformComponent.WorldPosition:ToVector2(), self.Entity.TransformComponent.WorldPosition:ToVector2())
		)
	
	self.Entity.LineRendererComponent.Enable = true
}

[server]
void Disconnect ()
{
	self.Entity.DistanceJointComponent.Joints:Clear()
	
	self.Entity.LineRendererComponent.Enable = false
}

Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.R then
		if self.Entity.DistanceJointComponent:GetJointsCount() == 0 then
			self:Connect()
		else
			self:Disconnect()
		end
	end
}
```

# SeeAlso

- [LineRendererComponent](https://mod-developers.nexon.com/apiReference/Components/LineRendererComponent)
- [TransformComponent](https://mod-developers.nexon.com/apiReference/Components/TransformComponent)
- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-10-28 PM 02:21


# FootholdComponent

맵의 모든 발판을 관리합니다. 발판은 RigidbodyComponent와 상호작용합니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetFoothold(int32 footholdId) |
| --- |
| footholdId에 해당하는 Foothold를 반환합니다. |

| [table<Foothold>](https://mod-developers.nexon.com/apiReference/Lua/table) GetFootholdAll() |
| --- |
| 맵상의 전체 Foothold들을 반환합니다. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetNearestFootholdByPoint([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, float distance) |
| --- |
| 입력한 점에서 distance 범위 안의 가장 가까운 Foothold를 찾아 반환합니다. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) Raycast([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float distance) |
| --- |
| point 위치에서 direction 방향으로 distance 만큼 이동하는 광선을 쏘고, 이 광선과 충돌하는 첫번째 Foothold를 찾아 반환합니다. |

| [table<Foothold>](https://mod-developers.nexon.com/apiReference/Lua/table) RaycastAll([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float distance) |
| --- |
| point 위치에서 direction 방향으로 distance 만큼 이동하는 광선을 쏘고, 이 광선과 충돌하는 모든 Foothold들을 찾아 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

화면을 터치하면 랜덤한 한 수를 뽑아 `GetNearestFootholdByPoint`, `Raycast`, `RaycastAll`를 매칭합니다. 이 세 함수를 사용해 탐색되는 동적 발판을 파괴하는 예제입니다.

```
Method:
[server]
void GetNearstFootholdByPoint(Vector point, number distance)
{
	local footholdComponent = self.Entity.CurrentMap.FootholdComponent
	if footholdComponent == nil then
		return
	end
	 
	local foothold = footholdComponent:GetNearestFootholdByPoint(point, distance)
	if foothold == nil then
		log("There is no close foothold.")
		return
	end
	 
	local entity = _EntityService:GetEntity(foothold.OwnerId)
	if entity == nil or entity.CustomFootholdComponent == nil then
		return
	end
	 
	if entity.CustomFootholdComponent.IsDynamicFoothold == false then
		return
	end
	 
	entity:Destroy()
}

[server]
void Raycast(Vector direct, number distance)
{
	local footholdComponent = self.Entity.CurrentMap.FootholdComponent
	if footholdComponent == nil then
		return
	end
	 
	local foothold = footholdComponent:Raycast(self.Entity.TransformComponent.Position:ToVector2(), direct:Normalize(), distance)
	if foothold == nil then
		log("There is no close foothold.")
		return
	end
	 
	local entity = _EntityService:GetEntity(foothold.OwnerId)
	if entity == nil or entity.CustomFootholdComponent == nil then
		return
	end
	 
	if entity.CustomFootholdComponent.IsDynamicFoothold == false then
		return
	end
	 
	entity:Destroy()
}

[server]
void RaycastAll(Vector2 direct, number distance)
{
	local footholdComponent = self.Entity.CurrentMap.FootholdComponent
	if footholdComponent == nil then
		return
	end
	 
	local footholds = footholdComponent:RaycastAll(self.Entity.TransformComponent.Position:ToVector2(), direct:Normalize(), distance)
	if footholds == nil then
		log("There is no close foothold.")
		return
	end
	 
	for i, j in ipairs(footholds) do
		local entity = _EntityService:GetEntity(j.OwnerId)
		if entity == nil or entity.CustomFootholdComponent == nil then
			return
		end
		 
		if entity.CustomFootholdComponent.IsDynamicFoothold == false then
			return
		end
		 
		entity:Destroy()
	end
}

Event Handler:
[service: InputService]
HandleScreenTouchEvent(ScreenTouchEvent event)
{
	-- Parameters
	local TouchPoint = event.TouchPoint
	--------------------------------------------------------
	local rand = _UtilLogic:RandomIntegerRange(1,3)
	log(rand)
	
	if rand == 1 then
		self:GetNearestFootholdByPoint(_UILogic:ScreenToWorldPosition(TouchPoint), 1)
	elseif rand == 2 then
		self:Raycast(_UILogic:ScreenToWorldPosition(TouchPoint), 5)
	else
		self:RaycastAll(_UILogic:ScreenToWorldPosition(TouchPoint), 10)
	end
}
```

# SeeAlso

- [CustomFootholdComponent](https://mod-developers.nexon.com/apiReference/Components/CustomFootholdComponent)
- [UILogic](https://mod-developers.nexon.com/apiReference/Logics/UILogic)
- [math](https://mod-developers.nexon.com/apiReference/Lua/math)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [발판 만들기](/docs?postId=71)
- [이동 발판 만들기](/docs?postId=579)

Update 2025-10-28 PM 02:21


# GridViewComponent

규격화된 UI Entity들을 Grid 형태로 표현합니다. 화면에 보이는 UI Entity만 생성하고 재사용합니다. 대량의 Grid를 표현하기에 최적화된 형태입니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CellSize |
| --- |
| 자식 UI Entity의 고정 크기입니다. |

| int32 FixedCount |
| --- |
| 고정된 행 혹은 열의 개수를 설정합니다. |

| [GridViewFixedType](https://mod-developers.nexon.com/apiReference/Enums/GridViewFixedType) FixedType |
| --- |
| 행을 고정할지 열을 고정할지 설정합니다. |

| [HorizontalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/HorizontalScrollBarDirection) HorizontalScrollBarDirection |
| --- |
| 가로 스크롤바의 방향입니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) ItemEntity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 자식 UI Entity 생성 시 복제될 원본입니다. ItemEntity로 설정된 엔티티는 disable 됩니다. |

| function<int32, Entity> OnClear ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 자식 UI Entity가 사용되지 않을 때 호출되는 콜백입니다. 자식 UI Entity의 인덱스와 UI Entity가 전달됩니다. |

| function<int32, Entity> OnRefresh ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 자식 UI Entity가 재사용될 때 호출되는 콜백입니다. 자식 UI Entity 의 인덱스와 UI Entity가 전달됩니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| 그리드뷰의 상하좌우 여유 공간을 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarBackgroundColor |
| --- |
| 스크롤바의 배경 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarBackgroundImageRUID |
| --- |
| 스크롤바의 배경 이미지입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarHandleColor |
| --- |
| 스크롤바의 핸들 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarHandleImageRUID |
| --- |
| 스크롤바의 핸들 이미지입니다. |

| float ScrollBarThickness |
| --- |
| 스크롤바 영역의 두께입니다. |

| [ScrollBarVisibility](https://mod-developers.nexon.com/apiReference/Enums/ScrollBarVisibility) ScrollBarVisible |
| --- |
| 스크롤바의 표시 여부를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Spacing |
| --- |
| 자식 UI Entity간의 간격입니다. |

| int32 TotalCount |
| --- |
| GridView에 표시될 자식 UI Entity의 총개수입니다. |

| boolean UseScroll |
| --- |
| 스크롤 기능 사용 여부를 설정합니다. |

| [VerticalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/VerticalScrollBarDirection) VerticalScrollBarDirection |
| --- |
| 세로 스크롤바의 방향입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetScrollNormalizedPosition() |
| --- |
| 스크롤바의 정규화된 위치를 반환합니다. |

| float GetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| 지정한 방향 스크롤바의 정규화된 위치를 반환합니다. |

| void Refresh(boolean resetPos = true, boolean force = false) |
| --- |
| GridView의 자식 UI Entity를 전체 갱신합니다. resetPos가 true라면 갱신 후 스크롤 위치를 초기화합니다. |

| void RefreshIndex(int32 index) |
| --- |
| 특정 index의 자식 UI Entity를 갱신합니다. 자식 UI Entity의 index와 UI Entity가 OnRefresh 콜백을 통해 호출됩니다. |

| void ResetScrollPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| 지정한 축의 스크롤바 위치를 처음 위치인 위쪽으로 이동합니다. |

| void SetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis, float value) |
| --- |
| 지정한 축의 스크롤바 위치를 지정한 정규화된 위치로 이동합니다. 위쪽이 0, 아래쪽이 1입니다. |

| void SetScrollPositionByItemIndex(int32 index) |
| --- |
| 스크롤바 위치를 특정 index의 자식 UI Entity가 보이는 위치로 이동합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ScrollPositionChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ScrollPositionChangedEvent) |
| --- |
| 스크롤이 가능한 UI Entity에서 스크롤 위치가 변경될 때 발생하는 이벤트입니다. UI Entity에 ScrollLayoutGroupComponent 또는 GridViewComponent가 있어야 이벤트가 발생합니다. |

# Examples

레이아웃에 따라 엔티티를 정렬하는 예제입니다.

```
Property:
[None]
GridViewComponent Inventory = nil
[None]
Entity InventorySlot = nil
  
Method:
[client only]
void OnBeginPlay ()
{
	self.Inventory.ItemEntity = self.InventorySlot
	self.Inventory.OnRefresh = self.OnRefresh
	self.Inventory.OnClear = self.OnClear
	self.Inventory.TotalCount = 10
	self.Inventory:Refresh()
}
 
void OnRefresh (number index, Entity entity)
{
	entity:GetChildByName("Name").TextComponent.Text = tostring(index)
}
 
void OnClear (number index, Entity entity)
{
	entity:GetChildByName("Name").TextComponent.Text = ""
}
```

# SeeAlso

- [TextComponent](https://mod-developers.nexon.com/apiReference/Components/TextComponent)
- [global](https://mod-developers.nexon.com/apiReference/Lua/global)

Update 2025-08-27 PM 04:56


# HitComponent

엔티티의 충돌 영역을 설정하고, AttackComponent의 공격을 받았을 때 피격 동작을 구현할 수 있는 인터페이스를 제공합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| IsLegacy가 true인 이전 시스템에서 사용할 수 있습니다. 엔티티를 기준으로 충돌체 직사각형의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 직사각형 충돌체의 너비와 높이를 지정합니다. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 원형 충돌체의 반지름입니다. ColliderType이 Circle일 때 유효합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. CollisionGroup을 사용하세요. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Entity를 기준으로 충돌체의 중심점 위치를 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌체의 타입을 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌 그룹을 설정합니다. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 Component가 이전의 시스템으로 동작하는지 여부입니다. 신규 시스템은 충돌체가 TransformComponent의 회전과 크기의 적용을 받습니다. 또한 ColliderType을 설정해 원 모양 충돌체를 사용할 수 있습니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형 충돌체를 이루는 점들의 위치입니다. ColliderType이 Polygon일 때 유효합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| boolean IsHitTarget([string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 엔티티가 AttackComponent의 공격을 받을지 말지 여부를 판단합니다. 기본 동작으로 true를 반환합니다.<br>attackInfo는 AttackComponent로부터 전달된 사용자 정의 데이터입니다. |

| void OnHit([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, integer damage, boolean isCritical, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, int32 hitCount) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 엔티티가 피격되었을 때 호출됩니다. 기본 동작으로 HitEvent를 발생시킵니다.<br>attacker는 공격한 Entity, attackInfo는 AttackComponent로부터 전달된 사용자 정의 데이터, hitCount는 대미지 분할 재생 횟수입니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [HitEvent](https://mod-developers.nexon.com/apiReference/Events/HitEvent) |
| --- |
| 엔티티가 피격되었을 때 발생하는 이벤트입니다. |

# Examples

다음은 몬스터가 플레이어의 공격을 받을 때마다 `HitComponent`의 BoxSize를 남은 체력에 반비례하게 키우는 예제입니다.

```
Property:
[Sync]
number Health = 1000
[None]
number InitialHealth = 0
[None]
Vector2 InitialBoxSize = Vector2(0,0)
  
Method:
[server only]
void OnBeginPlay ()
{
	self.InitialHealth = self.Health
	self.InitialBoxSize = self.Entity.HitComponent.BoxSize:Clone()
}
  
Entity Event Handler:
[server only] [self]
HandleHitEvent (HitEvent event)
{
	-- Parameters
	local AttackCenter = event.AttackCenter
	local AttackerEntity = event.AttackerEntity
	local Damages = event.Damages
	local Extra = event.Extra
	local FeedbackAction = event.FeedbackAction
	local IsCritical = event.IsCritical
	local TotalDamage = event.TotalDamage
	--------------------------------------------------------
	  
	local hitComponent = self.Entity.HitComponent
	  
	self.Health = self.Health - TotalDamage
	self.Health = math.max(self.Health, 0.0)
	  
	if self.Health > 0.0 then
		local ratio = 10 - ((10 - 1) / self.InitialHealth) * self.Health
		hitComponent.BoxSize = self.InitialBoxSize * ratio
	else
		_EntityService:Destroy(self.Entity)
	end
}
```

`HitComponent`를 확장해 피격 여부 판정과 피격 시 동작을 직접 구현할 수 있습니다. 다음은 기본 제공 스크립트인 PlayerHit 스크립트의 IsHitTarget 함수입니다. 피격 시 1초간 플레이어를 무적 상태로 만듭니다.

```
Property:
[None]
number ImmuneCooldown = 1
[None]
number LastHitTime = 0
  
Method:
override boolean IsHitTarget()
{
	local currentTime = _UtilLogic.ElapsedSeconds
		if self.LastHitTime + self.ImmuneCooldown < currentTime then
		self.LastHitTime = _UtilLogic.ElapsedSeconds
	return true
	end
	
	return false
}
```

# SeeAlso

- [AttackComponent](https://mod-developers.nexon.com/apiReference/Components/AttackComponent)
- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [공격과 피격](/docs?postId=206)

Update 2025-08-27 PM 04:56


# HitEffectSpawnerComponent

엔티티에 피격 이벤트가 발생했을 때 피격 이펙트를 출력합니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# ImageComponent

SpriteGUIRendererComponent등과 같이 UI에 이미지를 출력하는 컴포넌트의 부모 컴포넌트입니다. 엔티티에 직접 추가해서 사용할 수 없습니다.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 이미지의 기본 색상을 설정합니다 |

| boolean DropShadow |
| --- |
| 이미지의 그림자 출력 여부를 설정합니다. |

| float DropShadowAngle |
| --- |
| 그림자를 출력할 각도를 설정합니다 |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| 그림자 색상을 설정합니다. |

| float DropShadowDistance |
| --- |
| 이미지와 그림자의 거리입니다. |

| float FillAmount |
| --- |
| Type이 Filled로 설정되어 있을 때 표시되는 이미지의 비율입니다. 0부터 1 사이의 값을 사용합니다. |

| boolean FillCenter |
| --- |
| Type이 Sliced 또는 Tiled로 설정되어 있을 때 이미지 영역의 가운데를 채울지를 설정합니다. |

| boolean FillClockWise |
| --- |
| FillMethod가 Radial90, Radial180, Radial360으로 설정되어 있을 때 채우기 방향을 설정합니다. 값이 true면 시계 방향으로 채워집니다. |

| [FillMethodType](https://mod-developers.nexon.com/apiReference/Enums/FillMethodType) FillMethod |
| --- |
| Type이 Filled일 때의 채우기 방식을 설정합니다. |

| int32 FillOrigin |
| --- |
| Type이 Filled로 설정되어 있을 때 채우기 시작점을 설정합니다. FillMethod가 Horizontal 또는 Vertical일 경우 0 ~ 1 값을 사용할 수 있습니다. Radial90, Radial180, Radial360일 경우 0 ~ 3 값을 사용할 수 있습니다. |

| boolean FlipX |
| --- |
| 이미지의 X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY |
| --- |
| 이미지의 Y축을 기준으로 반전 여부를 결정합니다. |

| int32 FrameColumn ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| int32 FrameRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| int32 FrameRow ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ImageRUID |
| --- |
| 화면에 표시될 이미지 RUID입니다. |

| boolean Outline |
| --- |
| 이미지 외곽선 출력 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| 이미지 외곽선 색상입니다. |

| float OutlineWidth |
| --- |
| 외곽선 두께입니다. |

| boolean RaycastTarget |
| --- |
| true로 설정할 경우 화면 터치 또는 마우스 클릭 대상이 되며, 뒤에 가려진 UI는 화면 터치와 마우스 클릭 입력을 받지 못합니다. |

| [ImageType](https://mod-developers.nexon.com/apiReference/Enums/ImageType) Type |
| --- |
| 이미지를 표시하는 방식입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void SetNativeSize() |
| --- |
| 이미지를 원본 크기로 조정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# InteractionComponent

플레이어가 엔티티와 상호작용할 수 있는 기능을 제공합니다.

# Properties

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) ActionKey ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Interaction에 사용할 Key를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ActionName ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 말풍선에 표시될 상호작용의 이름을 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| IsLegacy가 true인 이전 시스템에서 사용할 수 있습니다. 엔티티를 기준으로 충돌체 직사각형의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 직사각형 충돌체의 너비와 높이를 지정합니다. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 원형 충돌체의 반지름입니다. ColliderType이 Circle일 때 유효합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. CollisionGroup을 사용하세요. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티를 기준으로 충돌체의 중심점 위치를 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌체의 타입을 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 충돌체의 충돌 그룹입니다. |

| float HoldingDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| InteractionType이 KeyHoldingDuration 또는 KeyUpAfterHoldingDuration일 때 얼마 동안 키를 누르고 있어야 하는지 설정합니다. 초 단위를 사용합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) InfoOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [InteractType](https://mod-developers.nexon.com/apiReference/Enums/InteractType) InteractionType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 상호작용을 위해 필요한 키 입력 방식을 설정합니다. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 Component가 이전의 시스템으로 동작할지를 설정합니다. 신규 시스템은 충돌체가 TransformComponent의 회전과 크기에 영향을 받습니다. 또한 ColliderType을 설정해 원 모양 충돌체를 사용할 수 있습니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형 충돌체를 이루는 점들의 위치입니다. ColliderType이 Polygon일 때 유효합니다. |

| boolean ShowActionInfo ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 플레이어가 가까이 접근했을 때 말풍선이 나타나 ActionName과 ActionKey를 보여줄지 여부를 설정합니다. 말풍선을 표시하기 위해 플레이 중 자동으로 ChatBalloonComponent가 추가됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void OnEnter() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionEnterEvent를 활용하세요. |

| void OnInteraction() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionEvent를 활용하세요. |

| void OnLeave() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionLeaveEvent를 활용하세요. |

| void SetOnEnter(func onEnterFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionEnterEvent를 활용하세요. |

| void SetOnInteraction(func onInteractionFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionEvent를 활용하세요. |

| void SetOnLeave(func onLeaveFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. InteractionLeaveEvent를 활용하세요. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [InteractionEnterEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEnterEvent) |
| --- |
| Interaction영역 안으로 들어올 때 발생하는 이벤트입니다. |

| [InteractionEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEvent) |
| --- |
| Interaction 했을 때 발생하는 이벤트입니다. |

| [InteractionLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionLeaveEvent) |
| --- |
| Interaction 영역 밖으로 벗어날 때 발생하는 이벤트입니다. |

# Examples

Interaction 될 경우 Player Entity가 가지고 있는 MovementComponent의 InputSpeed가 3으로 변화합니다.

```
Event Handler:
[self]
HandleInteractionEnterEvent (InteractionEnterEvent event)
{
	-- Parameters
	local InteractionEntity = event.InteractionEntity
	--------------------------------------------------------
	if self:IsClient() then
		return
	end
 
	local effectRoot = _EntityService:GetEntityByPath(EntityPath)
	effectRoot.Enable = true
}

[self]
HandleInteractionLeaveEvent (InteractionLeaveEvent event)
{
	-- Parameters
	local InteractionEntity = event.InteractionEntity
	--------------------------------------------------------
	if self:IsClient() then
		return
	end
	 
	local effectRoot = _EntityService:GetEntityByPath(EntityPath)
	effectRoot.Enable = false
}

[self]
HandleInteractionEvent (InteractionEvent event)
{
	-- Parameters
	local InteractionEntity = event.InteractionEntity
	--------------------------------------------------------
	if self:IsServer() then
		return
	end
	 
	InteractionEntity.MovementComponent.InputSpeed = 3
}
```

# SeeAlso

- [MovementComponent](https://mod-developers.nexon.com/apiReference/Components/MovementComponent)
- [InteractionEnterEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEnterEvent)
- [InteractionEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEvent)
- [InteractionLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionLeaveEvent)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)

Update 2025-08-27 PM 04:56


# InventoryComponent

소유한 아이템들을 관리하는 컴포넌트입니다.

# Properties

| boolean IsInitialized ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 인벤토리의 초기화가 완료 여부를 확인합니다. 초기화된 이후 InventoryComponent 관련 기능이 정상 작동하게 됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [table<Item>](https://mod-developers.nexon.com/apiReference/Lua/table) GetItemList() |
| --- |
| 소유한 아이템들을 가져옵니다. |

| [table<Item>](https://mod-developers.nexon.com/apiReference/Lua/table) GetItemsWithType(Type itemType) |
| --- |
| 소유한 아이템들 중 입력한 타입의 아이템을 가져옵니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [InventoryItemAddedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemAddedEvent) |
| --- |
| 인벤토리의 아이템이 추가되었을 때 발생하는 이벤트입니다. |

| [InventoryItemInitEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemInitEvent) |
| --- |
| 인벤토리가 초기화되었을 때 발생하는 이벤트입니다. 이 이벤트가 발생한 이후 InventoryComponent 관련 기능이 정상 작동하게 됩니다. |

| [InventoryItemModifiedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemModifiedEvent) |
| --- |
| 인벤토리의 아이템이 수정되었을 때 발생하는 이벤트입니다. |

| [InventoryItemRemovedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemRemovedEvent) |
| --- |
| 인벤토리의 아이템이 제거되었을 때 발생하는 이벤트입니다. |

# Examples

인벤토리에 아이템이 추가, 제거되거나 변경 사항이 생기면 `InventoryComponent`는 해당하는 이벤트를 발생시킵니다. `ItemService`의 예제 코드를 예시로 각 이벤트의 호출 시점을 살펴보면 다음과 같습니다.

```
Event Handler:
[self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	if self:IsClient() then
	    return
	end
	 
	local inventory = self.Entity.InventoryComponent
	local items = inventory:GetItemList()
	 
	if TriggerBodyEntity.Name == "Get Item" then
	    local newItem = _ItemService:CreateItem(TestItem, "Test Item", inventory)
	    newItem.ItemCount = 3
	elseif TriggerBodyEntity.Name == "Give Item" then
	    if #items > 0 then
	        items[1].ItemCount = items[1].ItemCount - 1
	        if items[1].ItemCount == 0 then
	            _ItemService:RemoveItem(items[1])
	        end
	    end
	elseif TriggerBodyEntity.Name == "Trash Can" then
	    if #items > 0 then
	        _ItemService:RemoveItem(items[1])
	    end
	end
}
```

`_ItemService:CreateItem` 함수와 `_ItemService:RemoveItem` 함수로 인해 인벤토리에 아이템이 추가, 제거될 때 각각 `InventoryItemAddedEvent`, `InventoryItemRemovedEvent` 이벤트가 발생합니다.

```
Event Handler:
[self]
HandleInventoryItemAddedEvent (InventoryItemAddedEvent event)
{
	-- Parameters
	local Entity = event.Entity
	local Items = event.Items
	--------------------------------------------------------
	if self:IsServer() then
	    return
	end
	 
	log("[Get Item]")
}

[self]
HandleInventoryItemRemovedEvent (InventoryItemRemovedEvent event)
{
	-- Parameters
	local Entity = event.Entity
	local Items = event.Items
	--------------------------------------------------------
	if self:IsServer() then
	    return
	end
	 
	log("[Item Removed]")
}
```

ItemCount 프로퍼티 값이 변할 때 `InventoryItemModifiedEvent` 이벤트가 발생합니다. 그 외 이벤트가 발생하는 경우는 아래와 같습니다.

- 아이템의 소유자(Owner)가 변경될 때
- 아이템 타입의 프로퍼티 값이 변경될 때

```
Event Handler:
[self]
HandleInventoryItemModifiedEvent (InventoryItemModifiedEvent event)
{
	-- Parameters
	local Entity = event.Entity
	local Items = event.Items
	--------------------------------------------------------
	if self:IsServer() then
	    return
	end
	 
	log ("[Give Item]")
}
```

인벤토리는 게임을 나가도 계속 유지됩니다. 게임에 입장하면 발생하는 `InventoryItemInitEvent` 이벤트에서 아이템 목록을 확인합니다.

```
Event Handler:
[self]
HandleInventoryItemInitEvent (InventoryItemInitEvent event)
{
	-- Parameters
	local Entity = event.Entity
	local Items = event.Items
	--------------------------------------------------------
	if self:IsServer() then
	    return
	end
	 
	log("[Current Items]")
	for i = 1, #Items do
	    log("Item" .. i .. ": ", Items[i].ItemDataTableName)
	end	
}
```

# SeeAlso

- [InventoryItemAddedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemAddedEvent)
- [InventoryItemInitEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemInitEvent)
- [InventoryItemModifiedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemModifiedEvent)
- [InventoryItemRemovedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemRemovedEvent)
- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [ItemService](https://mod-developers.nexon.com/apiReference/Services/ItemService)

Update 2025-08-27 PM 04:56


# JoystickComponent

모바일 환경에서 플레이어의 이동을 제어할 수 있는 가상의 조작키 기능을 지원해주는 Component입니다.

# Properties

| [AxisType](https://mod-developers.nexon.com/apiReference/Enums/AxisType) Axis |
| --- |
| 조작키가 움직일 수 있는 축 타입을 설정합니다. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) DownArrow |
| --- |
| 조작키가 아래로 이동했을 때 입력되는 키입니다. |

| boolean DynamicStick |
| --- |
| 플레이 시 조작키의 위치가 유저가 터치한 곳으로 이동합니다. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) LeftArrow |
| --- |
| 조작키가 왼쪽으로 이동했을 때 입력되는 키입니다. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) RightArrow |
| --- |
| 조작키가 오른쪽로 이동했을 때 입력되는 키입니다. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) UpArrow |
| --- |
| 조작키가 위로 이동했을 때 입력되는 키입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

조이스틱을 아래로 당겼을 때 키보드 E키를 누른 것럼 동작하게 설정하는 예제입니다. 이동은 클라이언트 공간에서 동작하므로 키 또한 클라이언트에서 변경되어야 합니다.

```
Method:
[client]
void SetJoystickDownArrow ()
{
	local joystick = _EntityService:GetEntityByPath("/ui/DefaultGroup/UIJoystick")
	joystick.JoystickComponent.DownArrow = KeyboardKey.E
}

Event Handler:
[client only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	self:SetJoystickDownArrow()
}
```

# SeeAlso

- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)

Update 2025-08-27 PM 04:56


# KinematicbodyComponent

탑다운 방식의 상하좌우 이동, 점프 및 렉트 타일과의 충돌 기능을 제공합니다. 중력과 가·감속 영항을 받지 않습니다. 타일맵이 RectTile일 때 동작합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Acceleration ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. SpeedFactor 프로퍼티를 사용하세요. |

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true인 경우 회전하거나 기울어진 사다리를 탄 캐릭터는 사다리의 모습을 따릅니다. false인 경우 캐릭터는 사다리의 기울기, 회전에 영향을 받지 않습니다. |

| boolean EnableJump ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 기능을 사용할지 여부를 나타냅니다. |

| boolean EnableShadow ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자를 사용할지 여부를 나타냅니다. |

| boolean EnableTileCollision ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| RectTileMap과의 충돌 기능을 사용할지 여부를 설정합니다. 값이 false일 경우 충돌 타일과 충돌하지 않습니다. 또한 RectTileCollisionBeginEvent 및 RectTileCollisionEndEvent도 발생하지 않습니다. |

| float JumpDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 속력 감소량입니다. 값이 클수록 지면에 더 빨리 떨어집니다. |

| float JumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 속력입니다. 값이 클수록 더 높이 점프합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 이동 속도를 설정합니다. SpeedFactor를 곱한 값이 최종 속도가 됩니다.<br>플레이어가 방향키를 이동하거나, MovementComponent:MoveToDirection() 함수를 호출했을 때 MoveVelocity 값이 변경됩니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ShadowColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자 색상입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ShadowOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자 위치입니다. |

| float ShadowScalingRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자 크기 변화율입니다. 엔티티의 점프 높이에 따라 크기가 변화합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ShadowSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자 크기입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이동할 때 X축, Y축 속력에 곱해지는 가중치입니다. 값이 클수록 이동 속력이 빨라집니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetGroundPosition() |
| --- |
| 로컬 좌표 기준으로 바닥 위치를 반환합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetWorldGroundPosition() |
| --- |
| 월드 좌표 기준으로 바닥 위치를 반환합니다. |

| boolean IsOnGround() |
| --- |
| 현재 지면에 닿아 있는 상태인지 확인합니다. 점프 중일 때 false를 반환합니다. |

| void OnEnterRectTile([RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) enterEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| RectTileEnterEvent 발생 시 호출되는 함수입니다. |

| void OnLeaveRectTile([RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) leaveEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| RectTileLeaveEvent 발생 시 호출되는 함수입니다. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 로컬 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 월드 좌표 기준으로 엔티티의 위치를 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [KinematicbodyJumpEvent](https://mod-developers.nexon.com/apiReference/Events/KinematicbodyJumpEvent) |
| --- |
| 점프 상태가 변경될 때 발생하는 이벤트입니다. |

| [RectTileCollisionBeginEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionBeginEvent) |
| --- |
| 충돌 가능한 타일과 접촉했을 때 발생하는 이벤트입니다. |

| [RectTileCollisionEndEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionEndEvent) |
| --- |
| 충돌한 타일에서 벗어날 때 발생하는 이벤트입니다. |

| [RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) |
| --- |
| 특정 사각형 타일에 진입했을 때 발생하는 이벤트입니다. |

| [RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) |
| --- |
| 특정 사각형 타일에서 벗어났을 때 발생하는 이벤트입니다. |

# Examples

점프로 타일을 뛰어 넘는 예제입니다.

```
Method:
[client only]
void OnUpdate ()
{
	if _UserService.LocalPlayer ~= self.Entity then
		return
	end
	 
	local kinematicbody = self.Entity.KinematicbodyComponent
	 
	local isOnGround = kinematicbody:IsOnGround()
	kinematicbody.EnableTileCollision = isOnGround
}
```

# SeeAlso

- [KinematicbodyJumpEvent](https://mod-developers.nexon.com/apiReference/Events/KinematicbodyJumpEvent)
- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)
- [RectTileMap에서 캐릭터 이동 제어](/docs?postId=748)

Update 2025-08-27 PM 04:56


# LineGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

UI에 선을 그리며, 선의 속성을 설정할 수 있는 기능을 제공합니다.

# Properties

| float Flexibility |
| --- |
| 꺾이는 부분의 너비를 조절합니다. IsFlexible이 true일 때 적용됩니다. |

| boolean IsFlexible |
| --- |
| 꺾이는 부분을 이어 그릴지 끊어 그릴지 설정합니다. |

| boolean IsSmooth |
| --- |
| true면 부드러운 곡선, false면 날카로운 꺾은 선으로 그립니다. |

| boolean Loop |
| --- |
| true면 자동으로 시작점과 끝점을 이어 닫힌 선을 그립니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| [SyncList<LinePoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points |
| --- |
| 선을 구성하는 정점들의 집합입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 랜더러에 적용할 머티리얼을 교체합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# LineRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

선을 그리며, 선의 속성을 설정할 수 있는 기능을 제공합니다.

# Properties

| float Flexibility ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 꺾이는 부분의 너비를 조절합니다. IsFlexible이 true일 때 적용됩니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsFlexible ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 꺾이는 부분을 이어 그릴지 끊어 그릴지 설정합니다. |

| boolean IsSmooth ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true면 부드러운 곡선, false면 날카로운 꺾은 선으로 그립니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true면 자동으로 시작점과 끝점을 이어 닫힌 선을 그립니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [SyncList<LinePoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 선을 구성하는 정점들의 집합입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 랜더러에 적용할 머티리얼을 교체합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

엔티티 주위에 원을 그리는 예제입니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	self:DrawCircle(2, 20, Color.red, 1)
}
 
[server]
void DrawCircle (number radius, integer vertexNum, Color color, number width)
{
	-- radius: 반지름
	-- vertexNum: 꼭지점 수
	-- color: 선 색상
	-- width: 선 두께
 
	if vertexNum < 3 then
		log_error ("The \'vertexNum\' should be greater than or equal to 3.")
		return
	end
	 
	local lineRenderer = self.Entity.LineRendererComponent
	lineRenderer.Points:Clear()
	lineRenderer.Loop = true
	 
	local delta = 360 / vertexNum
	 
	for i = 0, vertexNum - 1 do
		local theta = math.rad(90 + delta * i)
		local position = Vector2(math.cos(theta), math.sin(theta)) * radius
	 
		local point = LinePoint(position, color, width)
	 
		lineRenderer.Points:Add(point)
	end
}
```

# SeeAlso

- [math](https://mod-developers.nexon.com/apiReference/Lua/math)
- [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2)
- [다양한 선 그리기](/docs?postId=749)

Update 2025-10-28 PM 02:21


# MapComponent

맵을 관리하고, 맵의 고유 속성들을 변경할 수 있습니다.

# Properties

| float AirAccelerationXFactor |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티의 공중에서의 속도를 보정합니다. 값이 증가할수록 공중에서의 속도가 빨라집니다. |

| float AirDecelerationXFactor |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티가 공중에 있을 때, 입력이 없으면 X축 이동 속도가 얼마나 빠르게 멈추는지를 보정합니다. |

| float FallSpeedMaxXFactor |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티들의 공중에서의 X축 최대 속도 제한값을 보정합니다. |

| float FallSpeedMaxYFactor |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티의 공중에서의 Y축 최대 속도 제한값을 보정합니다. |

| float Gravity |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티의 중력값을 보정합니다. |

| boolean IsDynamicMap ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 동적으로 생성된 맵인지를 나타냅니다. |

| boolean IsInstanceMap |
| --- |
| 인스턴스 맵으로 사용할지 여부를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LeftBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 맵 영역의 좌하단 값입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RightTop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 맵 영역의 우상단 값입니다. |

| [TileMapMode](https://mod-developers.nexon.com/apiReference/Enums/TileMapMode) TileMapMode ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 맵의 타일맵 모드를 확인합니다. |

| boolean UseCustomBound ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 맵 영역을 직접 정의해 사용할지 여부를 나타냅니다. LeftBottom, RightTop 프로퍼티를 사용해 정의할 수 있습니다. false일 시 맵의 형태를 기반으로 자동 생성한 영역이 맵 영역으로 사용됩니다. TileMapMode가 MapleTile일 때에만 작동합니다. |

| float WalkAccelerationFactor |
| --- |
| 맵 내에 RigidbodyComponent를 가진 엔티티의 이동 속도의 Factor를 설정합니다. 최대 속도는 RigidbodyComponent의 WalkSpeed를 넘을 수 없으며, 최대 속도 도달까지의 속도를 제어합니다. |

| float WalkDrag |
| --- |
| 맵 내에 RigidbodyComponent를 가진 Entity들의 타일 마찰력을 보정합니다. 값이 작을수록 타일에서 잘 미끄러집니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| Vector2, Vector2 GetBound() |
| --- |
| LeftBottom, RightTop으로 구성된 맵 영역을 가져옵니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# MapLayerComponent

맵의 레이어별 정보를 관리합니다.

# Properties

| boolean IsVisible ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 레이어의 렌더링 여부를 나타냅니다. false일 경우 메이커 씬에서 보이지 않습니다. |

| int32 LayerSortOrder ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 레이어의 우선순위 순서입니다. 값이 작을수록 하위 레이어로, 뒤에 그려집니다. |

| boolean Locked ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 레이어가 잠겨 있는지를 나타냅니다. 잠겨 있으면 메이커에서 편집을 막습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MapLayerName ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 맵 레이어 이름입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# MaskComponent

마스크를 사용해, 자식 UI Entity의 특정 영역만 보이게 할 수 있습니다.

# Properties

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| 마스크의 상하좌우의 여유 공간을 설정합니다. |

| [MaskShape](https://mod-developers.nexon.com/apiReference/Enums/MaskShape) Shape |
| --- |
| 마스크의 모양을 지정합니다. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) Softness |
| --- |
| 마스크의 가장자리 영역부터 자연스럽게 번지게 할 영역을 정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

UI Entity 초기화 시 마스크에 패딩을 추가하여 보이는 영역을 축소합니다. `Softness`를 설정하여 자식 엔티티의 가장자리를 부드럽게 처리하는 예제입니다.

```
Method:
[client only]
void OnInitialize()
{
	self.Entity.MaskComponent.Softness = Vector2Int(10, 10)
	self.Entity.MaskComponent.Padding = RectOffset(10, 10, 10, 10)
}
```

# SeeAlso

- [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset)
- [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int)

Update 2025-08-27 PM 04:56


# MovementComponent

RigidbodyComponent, KinematicbodyComponent, SideviewbodyComponent를 제어하기 위한 이동 관련 기능을 제공합니다. 점프력과 속력을 간단하게 수정할 수 있습니다.

# Properties

| float InputSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이동 속력을 설정합니다. 값이 클수록 이동 속력이 빨라집니다.<br><br>- RigidbodyComponent가 true인 경우 좌우의 속력이 변경됩니다.<br>- KinematicMove가 true인 경우 상하좌우의 속력이 변경됩니다.<br>- KinamaticbodyComponent가 true인 경우 상하좌우 속력이 변경됩니다.<br>- SideviewbodyComponent가 true인 경우 좌우 속력이 변경됩니다. |

| boolean IsClimbPaused ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 등반 도중 멈춘 상태인지를 확인합니다. |

| float JumpForce ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 힘을 설정합니다. 값이 클수록 높게 점프합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| boolean DownJump() |
| --- |
| 아래로 점프합니다. 아래 점프 성공 여부를 반환합니다. |

| boolean IsFaceLeft() |
| --- |
| Entity가 왼쪽으로 향하는지 여부를 반환합니다. |

| boolean Jump() |
| --- |
| 점프합니다. 점프 성공 여부를 반환합니다. |

| void MoveToDirection([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float deltaTime) |
| --- |
| direction 방향으로 이동합니다. deltaTime은 초 단위이며, 사다리를 타고 있을 때만 적용됩니다. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 로컬 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 월드 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void Stop() |
| --- |
| 이동을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ChangedMovementInputEvent](https://mod-developers.nexon.com/apiReference/Events/ChangedMovementInputEvent) |
| --- |
| MovementComponent에서 이동 입력이 변경될 때 발생합니다. |

| [ClimbPauseEvent](https://mod-developers.nexon.com/apiReference/Events/ClimbPauseEvent) |
| --- |
| 물체를 타고 오르다 멈췄을 때 발생하는 이벤트입니다. |

# Examples

플레이어가 왼쪽을 바라보면 키 입력이 없어도 움직임이 시작됩니다, 또한 특정 엔티티에게 닿으면 점프 높이가 조절되고 점프하며 멈추는 행동을 하는 예제입니다.

```
[Sync]
boolean IsStarted = false
[Sync]
boolean IsFinished = false 

Method:
[client only]
void OnUpdate (number delta)
{
	if self.IsFinished then
		-- 입력에 의한 이동도 불가능
		self.Entity.MovementComponent:Stop() 
		return
	end
	 
	if self.IsStarted == false and self.Entity.MovementComponent:IsFaceLeft() then
		self.IsStarted = true
	end
 
	if self.IsStarted == false then
		return
	end
 
	self.Entity.MovementComponent:MoveToDirection(Vector2(1,0), delta)
}

Event Hanlder:
[client only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
 
	if TriggerBodyEntity.Name == "Name" then
		self.Entity.MovementComponent.JumpForce = 1.5
		self.Entity.MovementComponent:Jump()
	elseif TriggerBodyEntity.Name == "Name" then
		self.Entity.MovementComponent:DownJump()
	elseif TriggerBodyEntity.Name == "Name" then
		self.IsFinished = true
	end
}
```

# SeeAlso

- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [MovementComponent를 활용한 엔티티의 이동 제어](/docs?postId=546)

Update 2025-10-28 PM 02:21


# NameTagComponent

엔티티 이름표를 표시하고 관련 정보들을 설정합니다.

# Properties

| boolean Bold ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 굵은 텍스트 사용 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 글자색을 변경합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) FontOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이름표 내 글자 위치를 상하좌우로 조정합니다. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 글자 크기를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Name ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이름표에 표시할 이름을 설정합니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) NameTagRUID |
| --- |
| 이름표 형태를 변경 할 수 있습니다. |

| float OffsetY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이름표의 위치 Offset을 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

다음은 유저가 서로의 닉네임을 알 수 없도록 `NameTag`를 변경하는 예제입니다. 유저가 입장하면 UserId에 따라 정해진 단어와 색상으로 `NamgTag`를 설정합니다.

```
Event Handler:
[server only] [service: UserService]
HandleUserEnterEvent (UserEnterEvent event)
{
	-- Parameters
	local UserId = event.UserId
	--------------------------------------------------------
	 
	local userEntity = _UserService:GetUserEntityByUserId(UserId)
	local nametag = userEntity.NameTagComponent
	 
	if UserId == "000000" then
		nametag.Name = "Admin"
		nametag.FontColor = Color.magenta
	else
		nametag.Name = "Player"
		nametag.FontColor = Color.cyan
	end
}
```

# SeeAlso

- [UserEnterEvent](https://mod-developers.nexon.com/apiReference/Events/UserEnterEvent)
- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)
- [엔티티에 이름표 붙이기](/docs?postId=29)

Update 2025-08-27 PM 04:56


# PhysicsColliderComponent

물리 강체의 모양을 설정하고 다른 강체와 충돌 시 이벤트를 발생시킵니다. 자동으로 부모 또는 자신의 PhysicsRigidbodyComponent에 연결됩니다. 연결된 PhysicsRigidbodyComponent가 없으면 Static 강체로 취급됩니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Box Collider 크기를 설정합니다. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Circle Collider의 반지름 길이를 설정합니다. |

| boolean ClientOnly ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| true일 경우 Client에서 물리 연산이 일어나며, Client 공간에서 함수 사용 및 Property 변경이 가능합니다. 물리 연산 결과가 다른 Clients와 동기화 되지 않습니다. false일 경우 Server 공간에서 함수 사용 및 Property 변경이 가능합니다. 물리 연산 결과가 다른 Clients와 동기화 됩니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Collider의 오프셋을 설정합니다. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Collider의 모양을 설정합니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌 여부를 결정할 수 있는 충돌 그룹을 설정합니다. |

| float Density ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 연결된 PhysicsRigidbodyComponent의 UseDensity가 true일 경우 사용할 밀도 값을 설정합니다. |

| boolean EnableContactEvent ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| PhysicsContactBeginEvent와 PhysicsContactEndEvent 발생 여부를 설정합니다. false일 경우 두 Event가 발생하지 않습니다. |

| float Friction ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 마찰 계수를 설정합니다. 0 이상의 값만 설정할 수 있습니다. UseCustomPhysicalProperties의 값이 true일 경우 유효합니다. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Legacy 시스템을 지원할지를 설정합니다. 이전 시스템은 ColliderOffset이 Entity의 SpriteRendererComponent에 영향을 줍니다. 이전 시스템은 더 이상 지원하지 않으며, 추후 삭제 예정입니다. |

| boolean IsSensor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우 물리적인 상호작용이 일어나지 않지만, 충돌 이벤트는 발생합니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형 충돌체를 이루는 점들의 위치입니다. ColliderType이 Polygon일 때 유효합니다. |

| float Restitution ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌 후 튀어 오르는 정도를 설정합니다. 0 이상의 값만 설정할 수 있습니다. UseCustomPhysicalProperties가 true일 경우 유효합니다. |

| boolean UseCustomPhysicalProperties ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Friction과 Restituion 값이 물리 강체에 적용됩니다. false일 경우 연결된 PhysicsRigidbodyComponent의 값을 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [PhysicsRigidbodyComponent](https://mod-developers.nexon.com/apiReference/Components/PhysicsRigidbodyComponent) GetAttachedPhysicsRigidbody() |
| --- |
| 연결된 PhysicsRigidbodyComponent를 반환합니다. |

| void OnContactBegin([PhysicsContactBeginEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactBeginEvent) beginEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 두 물리 강체가 충돌을 시작할 때 호출됩니다. |

| void OnContactEnd([PhysicsContactEndEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactEndEvent) endEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 두 물리 강체의 충돌이 끝날 때 호출됩니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [PhysicsContactBeginEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactBeginEvent) |
| --- |
| 두 물리 강체가 충돌을 시작할 때 발생하는 이벤트입니다. |

| [PhysicsContactEndEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactEndEvent) |
| --- |
| 두 물리 강체가 충돌이 끝날 때 발생하는 이벤트입니다. |

# Examples

`P`를 눌렀을 때 CollisionGroup이 Monster로 변경되어 다른 Monster와 더이상 충돌하지 않게 만드는 예제입니다. `SleepingMode`를 `NeverSleep`으로 설정해야 합니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	self.Entity.PhysicsRigidbodyComponent:Sleep()
}

[server]
void Change ()
{
	self.Entity.SpriteRendererComponent.SpriteRUID = "000000"
	self.Entity.PhysicsColliderComponent.CollisionGroup = CollisionGroups.Monster
}

Event Handler:
[service: InputService]
HandleKeyDownEvent(KeyDownEvent event)
{
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.P then
	    self:Change()
	end
}
```

# SeeAlso

- [PhysicsRigidbodyComponent](https://mod-developers.nexon.com/apiReference/Components/PhysicsRigidbodyComponent)
- [SpriteRendererComponent](https://mod-developers.nexon.com/apiReference/Components/SpriteRendererComponent)
- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-12-02 PM 01:55


# PhysicsRigidbodyComponent

엔티티가 물리 엔진에 의해 제어됩니다. 물리 연산에 영향을 주는 값을 설정할 수 있습니다.

# Properties

| float AngularDamping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 각속도의 변화에 대한 저항 계수를 설정합니다. 값이 클수록 각속도 변화에 더 많은 힘이 필요합니다. 0 이상의 값만 설정할 수 있습니다. |

| [BodyType](https://mod-developers.nexon.com/apiReference/Enums/BodyType) BodyType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 강체 타입을 설정합니다. |

| [PhysicsCollisionDetectionMode](https://mod-developers.nexon.com/apiReference/Enums/PhysicsCollisionDetectionMode) CollisionDetectionMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 물리 충돌의 감지 방식을 설정합니다. |

| boolean FixedRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우 물체가 물리 상호작용에 의한 회전을 하지 않습니다. |

| float Friction ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 마찰 계수를 설정합니다. 값이 작을수록 잘 미끄러집니다. 0 이상의 값만 설정할 수 있습니다. |

| float GravityScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 중력에 영향을 받는 정도를 설정합니다. 값이 0일 경우 중력에 영향을 받지 않고, 1일 경우 중력 그대로 영향을 받습니다. 값이 커질수록 중력에 영향을 더 많이 받습니다. |

| float LinearDamping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 선형속도의 변화에 대한 저항 계수를 설정합니다. 값이 클수록 선형속도 변화에 더 많은 힘이 필요합니다. |

| float Mass ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 질량 값을 설정합니다. Collider의 넓이와 질량의 값을 이용해 밀도 값을 계산합니다 |

| float Restitution ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌 후 튀어 오르는 정도를 결정합니다. 0은 완전 비탄성 충돌, 1은 완전 탄성 충돌에 가깝습니다. |

| [PhysicsSleepingMode](https://mod-developers.nexon.com/apiReference/Enums/PhysicsSleepingMode) SleepingMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌의 수면 상태를 설정합니다. |

| boolean UseDensity ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우 mass의 값 대신 PhysicsColliderComponent의 Density와 Collider 값을 이용해 Mass 값을 계산합니다.<br>하나의 PhysicsRigidbodyComponent에 여러 개의 PhysicsColliderComponent가 붙으면 각 PhysicsColliderComponent의 Density와 Collider 크기를 곱한 값을 Mass 값으로 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ApplyAngularImpulse(float impulse) |
| --- |
| 강체에 반시계 방향으로(CCW) 각 충격량을 적용합니다. |

| void ApplyForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) force) |
| --- |
| 강체의 무게 중심에 특정 방향으로 힘을 가합니다. |

| void ApplyForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) force, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| 강체에 월드 상 특정 위치에서 특정 방향으로 힘을 가합니다. |

| void ApplyLinearImpulse([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) impulse) |
| --- |
| 강체의 무게 중심에 특정 방향으로 선형 충격량을 가합니다. |

| void ApplyLinearImpulse([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) impulse, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| 강체에 월드 상 특정 위치에서 특정 방향으로 선형 충격량을 가합니다. |

| void ApplyTorque(float force) |
| --- |
| 강체에 반시계 방향으로(CCW) 회전력을 가합니다. |

| void ClearPhysicsOwnership() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| 물리 소유권을 자동으로 부여합니다. Contact Event는 Server에서만 발생합니다. |

| float GetAngularVelocity() |
| --- |
| 강체의 각속도를 반환합니다. |

| float GetDensity() |
| --- |
| 강체의 밀도를 반환합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetLinearVelocity() |
| --- |
| 강체의 선형속도를 반환합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetLinearVelocityAtPoint([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| 월드 상 특정 위치에서의 강체 선형속도를 반환합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetPosition() |
| --- |
| 강체의 위치를 반환합니다. |

| float GetRotation() |
| --- |
| 강체의 회전각을 반환합니다. |

| void SetAngularVelocity(float velocity) |
| --- |
| 강체의 각속도를 설정합니다. |

| void SetClientAsPhysicsOwner([string](https://mod-developers.nexon.com/apiReference/Lua/string) userId) ![custom](https://img.shields.io/static/v1?label=&amp;message=ServerOnly&amp;color=mediumvioletred) |
| --- |
| 물리 소유권을 특정 유저에게 부여합니다. 해당 Client, Server에서 Contact Event가 발생합니다. |

| void SetLinearVelocity([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) velocity) |
| --- |
| 강체 선형속도를 설정합니다. 질량에 영향을 받지 않습니다. |

| void SetLinearVelocityX(float velocityX) |
| --- |
| 강체의 선형속도 X값을 설정합니다. 질량에 영향을 받지 않습니다. |

| void SetLinearVelocityY(float velocityY) |
| --- |
| 강체의 선형속도 Y값을 설정합니다. 질량에 영향을 받지 않습니다. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 강체의 위치를 설정합니다. |

| void SetRotation(float angle) |
| --- |
| 강체의 회전각을 반환합니다. |

| void SetServerAsPhysicsOwner() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| 물리 소유권을 서버에게 부여합니다. Server에서 Contact Event가 발생합니다. |

| void Sleep() |
| --- |
| 강체를 수면 상태로 전환합니다. |

| void Wake() |
| --- |
| 강체를 깨어 있는 상태로 전환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

#### ApplyLinearImpulse

```
Event Handler:
[self] 
HandlePhysicsContactBeginEvent (PhysicsContactBeginEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: PhysicsColliderComponent
	-- Space: Server, Client
	---------------------------------------------------------
	
	-- Parameters
	local ContactedBodyEntity = event.ContactedBodyEntity
	---------------------------------------------------------
	ContactedBodyEntity.PhysicsRigidbodyComponent:ApplyLinearImpulse(Vector2(0, 5))
}
```

#### GravityScale

```
Event Handler:
[self]
HandlePhysicsContactBeginEvent (PhysicsContactBeginEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: PhysicsColliderComponent
	-- Space: Server, Client
	---------------------------------------------------------
	
	-- Parameters
	local ContactedBodyEntity = event.ContactedBodyEntity
	---------------------------------------------------------
	ContactedBodyEntity.PhysicsRigidbodyComponent.GravityScale = 0.2
}
```

# SeeAlso

- [PhysicsContactBeginEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactBeginEvent)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-12-02 PM 01:55


# PhysicsSimulatorComponent

맵에 물리 법칙을 적용합니다. 물리 연산과 관련된 설정을 조절할 수 있습니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Gravity ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 물리 연산의 중력 값을 설정합니다. |

| boolean Paused ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 물리 연산의 일시정지 여부를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) WorldBounds ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 물리 연산 범위를 설정합니다. 엔티티가 범위를 벗어나면 물리 연산을 하지 않습니다. 예를 들어, 값을 (100, 100)으로 설정하면 원점(0, 0)을 중심으로 가로 200, 세로 200 크기의 사각형 영역이 물리 연산 범위가 됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void SetPositionIteration(int32 count) |
| --- |
| 물리 엔진의 위치 연산 반복 횟수를 설정합니다. 값이 클수록 물리 연산 결과가 정교해지며 더 많은 연산 시간이 소요됩니다. 최솟값은 1, 최댓값은 300입니다. 서버는 적용되지 않습니다. |

| void SetVelocityIteration(int32 count) |
| --- |
| 물리 엔진의 속도 연산 반복 횟수를 설정합니다. 값이 클수록 물리 연산 결과가 정교해지며 더 많은 연산 시간이 소요됩니다. 최솟값은 1, 최댓값은 30입니다. 서버는 적용되지 않습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

누르는 키에 따라 물리 시뮬레이션을 일시 정지시키거나, 중력 설정을 변경할 수 있는 예제입니다.

```
Event Handler:
[server only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	
	-- Parameters
	local inputValue = event.inputValue
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.P then
		self.Entity.CurrentMap.PhysicsSimulatorComponent.Paused = not self.Entity.CurrentMap.PhysicsSimulatorComponent.Paused
	elseif key == KeyboardKey.G then
		self.Entity.CurrentMap.PhysicsSimulatorComponent.Gravity = Vector2(0, 0)
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-10-28 PM 02:21


# PixelGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

픽셀 값을 지정해 원하는 스프라이트를 UI에 그릴 수 있는 기능을 제공합니다. 크기는 16x16 이하로 사용하는 것을 권장합니다.

# Properties

| int32 Height ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 스프라이트의 세로 길이입니다. |

| boolean IgnoreMapLayerCheck |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| int32 OrderInLayer |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| boolean RaycastTarget |
| --- |
| true로 설정할 경우 경우 화면 터치 또는 마우스 클릭 대상이 됩니다. 뒤에 가려진 UI는 입력을 받지 못합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| int32 Width ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 스프라이트의 가로 길이입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void FillColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트 전체의 색을 변경합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) GetPixel(int32 x, int32 y) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 원하는 위치의 픽셀 값을 반환합니다. 좌측 하단이 (1, 1)입니다. |

| [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixels() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 현재 스프라이트의 모든 픽셀 값을 반환합니다. 행 우선입니다.(Row-major) |

| [table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixelsAsRGBAInt() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정수로 나타낸 RGBA로 모든 픽셀 값을 반환합니다. 행 우선입니다.(Row-major) |

| void ResetWithColor(int32 width, int32 height, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트를 입력한 크기에 맞게 재설정하고, color 색으로 채웁니다. |

| void ResetWithColors(int32 width, int32 height, [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트를 입력한 크기에 맞게 재설정하고, 입력한 테이블의 값들로 픽셀들의 값을 설정합니다. 테이블 요소의 개수는 반드시 전체 픽셀 개수(Width*Height)와 같아야 합니다. 행 우선입니다.(Row-major) |

| void SetAlpha(float alpha) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트 전체의 알파 값을 변경합니다. |

| void SetPixel(int32 x, int32 y, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 입력한 위치의 픽셀 값을 설정합니다. 좌측 하단이 (1, 1)입니다. |

| void SetPixels([table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 입력한 테이블의 값들로 픽셀들의 값을 설정합니다. 테이블 요소의 개수는 반드시 전체 픽셀 개수(Width*Height)와 같아야 하고 행 우선입니다.(Row-major) |

| void SetPixelsByRGBAInt([table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정수로 나타낸 RGBA로 모든 픽셀 값을 설정합니다. 테이블의 길이는 반드시 전체 픽셀 개수(Width*Height)와 같아야 하고 행 우선입니다.(Row-major) |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

F 키를 누르면 픽셀이 무작위로 깜빡이는 예제입니다.

```
Property:
[None]
table Puzzle = {}
[None]
PixelGUIRendererComponent PixelGUIRendererComp
[None]
table PuzzlePos = {}

Method:
[client only]
void OnBeginPlay()
{
	self.PixelGUIRendererComp = self.Entity.PixelGUIRendererComponent
	self.PixelGUIRendererComp:FillColor(Color.black)
	  
	for i=1, 9 do
		local xPos = _UtilLogic:RandomIntegerRange(1, self.PixelGUIRendererComp.Width)
		local yPos = _UtilLogic:RandomIntegerRange(1, self.PixelGUIRendererComp.Height)
		self.Puzzles[i] = Vector2(xPos, yPos)
	end
}

[client only]
void StartPuzzle(number index)
{
	if index > #self.Puzzles then 
		return 
	end
	  
	self.PixelGUIRendererComp:SetPixel(self.Puzzles[index].x, self.Puzzles[index].y, Color.yellow)
	wait(0.35)
	self.PixelGUIRendererComp:SetPixel(self.Puzzles[index].x, self.Puzzles[index].y, Color.black)
	wait(0.1)
	self:StartPuzzle(index + 1)
}

EventHandler:
[service: InputService]
HandleKeyDownEvent(KeyDownEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	  
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.F then
		self:StartPuzzle(1)
	end
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)

Update 2025-12-02 PM 01:55


# PixelRendererComponent

원하는 스프라이트를 그릴 수 있는 기능을 제공합니다. 크기는 16x16 이하로 사용하는 것을 권장합니다.

# Properties

| int32 Height ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 스프라이트의 세로 길이입니다. |

| boolean IgnoreMapLayerCheck |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| int32 OrderInLayer |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| int32 Width ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 스프라이트의 가로 길이입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void FillColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트 전체를 입력한 색으로 채웁니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) GetPixel(int32 x, int32 y) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 입력한 위치의 픽셀 값을 반환합니다. 좌측 하단이 (1, 1)입니다. |

| [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixels() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 모든 픽셀 값을 반환합니다. 행 우선입니다.(Row-major) |

| [table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixelsAsRGBAInt() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정수로 나타낸 RGBA로 모든 픽셀 값을 반환합니다. 행 우선입니다.(Row-major) |

| void ResetWithColor(int32 width, int32 height, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트를 입력한 크기에 맞게 재설정하고, color 색으로 채웁니다. |

| void ResetWithColors(int32 width, int32 height, [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트를 입력한 크기에 맞게 재설정하고, 입력한 테이블의 값들로 픽셀들의 값을 설정합니다. 테이블 요소의 개수는 반드시 전체 픽셀 개수(Width*Height)와 같아야 합니다. 행 우선입니다.(Row-major) |

| void SetAlpha(float alpha) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 스프라이트의 알파 값을 설정합니다. |

| void SetPixel(int32 x, int32 y, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 입력한 위치의 픽셀 값을 설정합니다. 좌측 하단이 (1, 1)입니다. |

| void SetPixels([table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 입력한 테이블의 값들로 픽셀들의 값을 설정합니다. 테이블 요소의 개수는 반드시 전체 픽셀 개수(Width*Height)와 같아야 하고 행 우선입니다.(Row-major) |

| void SetPixelsByRGBAInt([table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정수로 나타낸 RGBA로 모든 픽셀 값을 설정합니다. 테이블의 길이는 반드시 전체 픽셀 개수(Width*Height)와 같아야 하고 행 우선입니다.(Row-major) |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

`F`키를 누르면 픽셀이 무작위로 깜빡이는 예제입니다.

```
Property:
[None]
table<Vector2> Puzzles
[None]
PixelRendererComponent PixelRendererComp = nil
[None]
table<Vector> PuzzlePos

Method:
[client only]
void OnBeginPlay ()
{
	self.PixelRendererComp = self.Entity.PixelRendererComponent
	self.PixelRendererComp:FillColor(Color.black)
	 
	for i=1, 9 do
		local xPos = _UtilLogic:RandomIntegerRange(1, 3)
		local yPos = _UtilLogic:RandomIntegerRange(1, 3)
		self.Puzzles[i] = Vector2(xPos, yPos)
	end
}

[client only]
void StartPuzzle (number index)
{
	if index > #self.Puzzles then return end
	 
	self.PixelRendererComp:SetPixel(self.Puzzles[index].x, self.Puzzles[index].y, Color.yellow)
	wait(0.35)
	self.PixelRendererComp:SetPixel(self.Puzzles[index].x, self.Puzzles[index].y, Color.black)
	wait(0.1)
	self:StartPuzzle(index + 1)
}

Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.F then
		self:StartPuzzle(1)
	end
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [스프라이트 색상을 픽셀 단위로 설정하기](/docs?postId=693)

Update 2025-12-02 PM 01:55


# PlayerComponent

플레이어를 나타내고, 관련 기능을 제공합니다.

# Properties

| integer Hp ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 Hp입니다. |

| integer MaxHp ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 최대 Hp입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Nickname ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 플레이어의 닉네임입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ProfileCode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 플레이어의 프로필 코드입니다. |

| boolean PVPMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 플레이어끼리 공격이 가능한지를 설정합니다. |

| float RespawnDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 죽은 후 리스폰까지 걸리는 시간입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) RespawnPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 리스폰될 위치를 설정합니다. 지정하지 않은 경우 1순위로 SpawnLocation의 위치, 2순위로 맵 진입 시점의 위치가 리스폰 위치로 설정됩니다. |

| number RespawnTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 리스폰이 될 예정 시간입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) UserId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 플레이어의 고유 식별자입니다. Client 실행제어 함수의 targetUserId 매개 변수에 사용할 수 있습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| boolean IsDead() |
| --- |
| 플레이어가 죽었는지 여부를 반환합니다. |

| void MoveToEntity([string](https://mod-developers.nexon.com/apiReference/Lua/string) entityID) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| entityID에 해당하는 엔티티와 같은 위치로 이동시킵니다. 다른 맵의 entityID인 경우 맵을 이동합니다. |

| void MoveToEntityByPath([string](https://mod-developers.nexon.com/apiReference/Lua/string) worldPath) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| worldPath에 존재하는 엔티티와 같은 위치로 이동시킵니다. worldPath가 다른 맵을 가리킬 경우 맵을 이동합니다. |

| void MoveToMapPosition([string](https://mod-developers.nexon.com/apiReference/Lua/string) mapID, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) targetPosition) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 플레이어를 특정 맵의 특정 위치로 이동시킵니다. |

| void ProcessDead([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 플레이어를 죽음에 이르게 합니다. |

| void ProcessRevive([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 플레이어를 부활시킵니다. |

| void Respawn() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 리스폰을 수행합니다. Respawn 위치의 1순위는 RespawnPosition, 2순위는 맵 내 SpawnLocation의 위치, 3순위는 맵 진입 시점의 위치입니다. |

| void SetPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) position) |
| --- |
| 로컬 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void SetWorldPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPosition) |
| --- |
| 월드 좌표 기준으로 엔티티의 위치를 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

"Check Point"라는 이름의 오브젝트에 닿으면 리스폰 위치를 오브젝트의 위치로 변경합니다. "Do Not Touch"라는 이름의 오브젝트에 닿으면 플레이어가 죽게 하여 리스폰 위치가 잘 변경됐는지 확인합니다.

```
[server only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event) 
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	 
	local player = self.Entity.PlayerComponent
	 
	if TriggerBodyEntity.Name == "Check Point" then
	    player.RespawnPosition = TriggerBodyEntity.TransformComponent.Position
	elseif TriggerBodyEntity.Name == "Do Not Touch" then
	    self.Entity.PlayerComponent:ProcessDead()
	end
}
```

# SeeAlso

- [TriggerComponent](https://mod-developers.nexon.com/apiReference/Components/TriggerComponent)
- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [플레이어 설정과 제어](/docs?postId=547)

Update 2025-10-28 PM 02:21


# PlayerControllerComponent

Player의 조작과 연관된 컴포넌트입니다. 입력과 액션을 연동하고 그 흐름을 제어합니다.

# Properties

| boolean AlwaysMovingState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 항상 걷기 애니메이션을 재생할지 여부를 결정합니다. |

| int32 FixedLookAt ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이동 시에 바라보는 방향을 한쪽으로 고정합니다. |

| float LookDirectionX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 현재 X축을 기준으로 캐릭터가 바라보고 있는 방향입니다. 양수일 경우 오른쪽, 음수일 경우 왼쪽입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ActionAttack() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Attack Key를 입력했을 때 동작입니다. |

| void ActionCrouch() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Crouch Key를 입력했을 때 동작입니다. |

| void ActionDownJump() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| ActionDownJump가 일어났을 때 동작입니다. |

| void ActionEnterPortal() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Portal Key를 입력했을 때 동작입니다. |

| void ActionInteraction([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key, boolean isKeyDown) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Interaction key를 입력했을 때 동작입니다. |

| void ActionJump() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Jump Key를 입력했을 때 동작입니다. |

| void ActionSit() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Sit Key를 입력했을 때 동작입니다. |

| void AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName, func -> boolean conditionFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| action 발동 조건을 추가합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetActionName([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| key에 매핑된 action의 이름을 반환합니다. |

| void RemoveActionKey([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정의된 key에 연결된 액션을 제거합니다. |

| void RemoveAllActionKeyByActionName([string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 정의된 이름에 연결된 액션을 모두 제거합니다. |

| void SetActionKey([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName, func -> boolean conditionFunction = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| key와 action을 매핑합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ChangedLookAtEvent](https://mod-developers.nexon.com/apiReference/Events/ChangedLookAtEvent) |
| --- |
| PlayerControllerComponent에서 캐릭터가 바라보는 방향이 변경될 때 발생합니다. |

| [PlayerActionEvent](https://mod-developers.nexon.com/apiReference/Events/PlayerActionEvent) |
| --- |
| 플레이어가 Action을 사용하면 발생하는 이벤트입니다. |

# Examples

"Jump", "Portal", "Crouch", "Attack", "Sit"와 같이 사전에 정의되어 있는 액션 이름에 키보드 입력을 추가로 연결하고, 사용자 정의 액션 이름을 키보드 입력과 연결하여 로그를 찍는 예제입니다.

키보드 B키로 공격, 키보드 N키로 점프가 가능해지며, 키보드 G키로 사용자 지정 액션 이름이 로그에 출력되는것을 확인할 수 있습니다.

```
Method:
[client only]
void OnBeginPlay ()
{
	self.Entity.PlayerControllerComponent:SetActionKey(KeyboardKey.B, "Attack")
	self.Entity.PlayerControllerComponent:SetActionKey(KeyboardKey.N, "Jump")	 
	self.Entity.PlayerControllerComponent:SetActionKey(KeyboardKey.G, "MyPlayerAction")
}

Event Handler:
[self]
HandlePlayerActionEvent (PlayerActionEvent event)
{
	-- Parameters
	local ActionName = event.ActionName
	local PlayerEntity = event.PlayerEntity
	--------------------------------------------------------
	log(ActionName)
}
```

# SeeAlso

- [AvatarRendererComponent](https://mod-developers.nexon.com/apiReference/Components/AvatarRendererComponent)
- [StateComponent](https://mod-developers.nexon.com/apiReference/Components/StateComponent)
- [StateStringToAvatarActionComponent](https://mod-developers.nexon.com/apiReference/Components/StateStringToAvatarActionComponent)
- [플레이어 설정과 제어](/docs?postId=547)

Update 2025-08-27 PM 04:56


# PolygonGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

UI에 다각형을 그리는 기능입니다.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 다각형의 색상입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points |
| --- |
| 다각형을 구성하는 꼭짓점들의 집합입니다. |

| boolean UseCustomUVs |
| --- |
| true라면 UVs에 설정한 UV 값이 도형에 적용됩니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) UVs |
| --- |
| 꼭짓점의 UV를 설정합니다. 반드시 Points와 길이가 같아야 합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 랜더러에 적용할 머티리얼을 교체합니다. |

| boolean IsDrawable() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Points 프로퍼티에 저장된 꼭짓점들을 이어 다각형이 그려지는지 확인합니다. 교차하는 변이 있으면 다각형이 그려지지 않습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# PolygonRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

다각형을 그리는 기능입니다.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형의 색상입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 사용할 머티리얼 Id를 지정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형을 구성하는 꼭짓점들의 집합입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| boolean UseCustomUVs ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true라면 UVs에 설정한 UV 값이 도형에 적용됩니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) UVs ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 꼭짓점의 UV를 설정합니다. 반드시 Points와 길이가 같아야 합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 랜더러에 적용할 머티리얼을 교체합니다. |

| boolean IsDrawable() |
| --- |
| Points 프로퍼티에 저장된 꼭짓점들을 이어 다각형이 그려지는지 확인합니다. 교차하는 변이 있으면 다각형이 그려지지 않습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# SeeAlso

- [다각형 그리기](/docs?postId=1080)

Update 2025-08-27 PM 04:56


# PortalComponent

포탈 기능을 제공합니다. 포탈의 충돌 속성을 설정하고 목적지 포탈을 연결할 수 있습니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티를 기준으로 충돌체 직사각형의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌체 직사각형의 너비와 높이를 지정합니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Portal의 충돌 그룹입니다. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트가 이전의 시스템으로 동작할지를 설정합니다. 신규 시스템은 충돌체가 TransformComponent의 회전과 크기에 영향을 받습니다. |

| [EntityRef](https://mod-developers.nexon.com/apiReference/Misc/EntityRef) PortalEntityRef |
| --- |
| 목적지 포탈을 설정합니다. PortalComponent를 가진 엔티티만 목적지로 설정할 수 있습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [PortalUseEvent](https://mod-developers.nexon.com/apiReference/Events/PortalUseEvent) |
| --- |
| 플레이어가 포탈을 이용할 때 발생하는 이벤트입니다. |

# Examples

맵에 3개 이상의 포탈을 배치합니다. 포탈에 컴포넌트를 추가하면 1초마다 연결된 포탈이 무작위로 변경됩니다.

```
Method:
[client only]
void OnBeginPlay ()
{
	-- 현재 맵의 모든 Entity를 순회하며 PortalComponent를 가진 Entity들의 EntityRef 저장
	self._T.portalRefList = {}
	local function findAndAddPortal(parent)
		-- 자기 자신은 제외
		if parent.PortalComponent ~= nil and parent ~= self.Entity then
			table.insert(self._T.portalRefList, EntityRef(parent))
		end
		if parent.Children ~= nil then
			for i, k in ipairs(parent.Children) do
				findAndAddPortal(k)
			end
		end
		end
	      
	findAndAddPortal(self.Entity.CurrentMap)
	      
	-- 자신 외에 포탈이 없다면 Return
	if #self._T.portalRefList < 1 then
		return
	end
	     
	-- 1초마다 연결된 포탈을 랜덤하게 변경하는 타이머 설정
	_TimerService:SetTimerRepeat(function()
		self.Entity.PortalComponent.PortalEntityRef = self._T.portalRefList[_UtilLogic:RandomIntegerRange(1, #self._T.portalRefList)]
	end, 1)
}
```

# SeeAlso

- [math](https://mod-developers.nexon.com/apiReference/Lua/math)
- [table](https://mod-developers.nexon.com/apiReference/Lua/table)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [다른 위치로 이동하는 포탈 만들기](/docs?postId=90)

Update 2025-08-27 PM 04:56


# PrismaticJointComponent

PrismaticJoint를 생성, 삭제할 수 있습니다. 연결된 강체가 특정 축 방향으로만 상대적 이동하도록 제한할 수 있습니다.

# Properties

| [SyncList<PrismaticJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Joint의 정보를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAxis) |
| --- |
| Joint를 추가합니다. 성공 시 index, 실패 시 -1을 반환합니다 |

| void DestroyJoint(int32 index) |
| --- |
| 순번이 index에 해당하는 Joint를 제거합니다. |

| int32 GetJointsCount() |
| --- |
| joint 수를 반환합니다. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| 순번이 index에 해당하는 Joint의 CollideConnected 값을 설정합니다. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorA 값을 설정합니다. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorB 값을 설정합니다. |

| void SetLocalAxis(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAxis) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAxis 값을 설정합니다. |

| void SetLowerTranslation(int32 index, float lowerTranslation) |
| --- |
| 순번이 index에 해당하는 Joint의 LowerTranslation 값을 설정합니다. |

| void SetMaxMotorForce(int32 index, float maxMotorForce) |
| --- |
| 순번이 index에 해당하는 Joint의 MaxMotorForce 값을 설정합니다. |

| void SetMotorEnable(int32 index, boolean enable) |
| --- |
| 순번이 index에 해당하는 Joint의 MotorEnable 값을 설정합니다. |

| void SetMotorSpeed(int32 index, float speed) |
| --- |
| 순번이 index에 해당하는 Joint의 MotorSpeed 값을 설정합니다. |

| void SetUpperTranslation(int32 index, float upperTranslation) |
| --- |
| 순번이 index에 해당하는 Joint의 UpperTranslation 값을 설정합니다. |

| void SetUseLimits(int32 index, boolean useLimits) |
| --- |
| 순번이 index에 해당하는 Joint의 UseLimits 값을 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

W, S를 누르면 SetMotorSpeed 값을 변경합니다. `PrismaticJointComponent`의 `UpperTranslation`, `LowerTranslation`의 값만큼만 TargetEntityRef로 설정한 Entity가 올라가고 내려갑니다.

#### SetMotorSpeed

```
Method:
[server]
void GoUp ()
{
	self.Entity.PrismaticJointComponent:SetMotorSpeed(1, 3)
}

[server]
void GoDown ()
{
	self.Entity.PrismaticJointComponent:SetMotorSpeed(1, -3)
}


Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.W then
		self:GoUp()
	elseif key == KeyboardKey.S then
		self:GoDown()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-10-28 PM 02:21


# PulleyJointComponent

PulleyJoint를 생성, 삭제합니다. 연결된 두 강체가 도르래처럼 이동합니다.

# Properties

| [SyncList<PulleyJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Joint 정보를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) groundAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) groundAnchorB, float ratio = 1) |
| --- |
| Joint를 추가합니다. 성공 시 index, 실패 시 -1을 반환합니다 |

| void DestroyJoint(int32 index) |
| --- |
| 순번이 index에 해당하는 Joint를 제거합니다. |

| int32 GetJointsCount() |
| --- |
| joint 수를 반환합니다. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| 순번이 index에 해당하는 joint의 CollideConnected 값을 설정합니다. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorA 값을 설정합니다. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorB 값을 설정합니다. |

| void SetRatio(int32 index, float ratio) |
| --- |
| 순번이 index에 해당하는 Joint의 Ratio 값을 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# SeeAlso

- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-10-28 PM 02:21


# RawImageGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

RawImage를 UI에 출력하는 기능을 제공합니다.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| void SetRawImage([RawImage](https://mod-developers.nexon.com/apiReference/Misc/RawImage) rawImage) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| RawImage를 그립니다. nil을 전달하면 아무것도 그리지 않습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# RawImageRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

RawImage를 출력하는 기능을 제공합니다.

# Properties

| boolean IgnoreMapLayerCheck |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼의 Id를 지정합니다. |

| int32 OrderInLayer |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

| void SetRawImage([RawImage](https://mod-developers.nexon.com/apiReference/Misc/RawImage) rawImage) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| RawImage를 그립니다. nil을 전달하면 아무것도 그리지 않습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# RectTileMapComponent

사각형 타일맵 기능을 제공합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GridSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 격자의 크기입니다. |

| boolean IgnoreMapLayerCheck |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsOddGridPosition |
| --- |
| 타일맵을 그리드의 기준점과 어긋나게 배치합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| true일 경우 Physics 기능을 사용하는 Dynamic 강체(PhysicRigidbody)와 충돌할 수 있습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 2개 이상의 Entity가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| int32 TileCount ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 배치되어 있는 타일의 총 개수입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) TileSetRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 타일셋의 RUID입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void BoxFill(int32 tileIndex, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| from부터 to까지의 사각형 영역에 타일을 배치합니다. tileIndex는 타일의 인덱스 번호입니다. |

| void BoxFill([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| from부터 to까지의 사각형 영역에 타일을 배치합니다. tileName은 타일의 Name 프로퍼티입니다. |

| void BoxRemove([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| from부터 to까지의 사각형 영역의 타일을 제거합니다. |

| void Clear() |
| --- |
| 타일들을 모두 제거합니다. |

| [List<Vector2Int>](https://mod-developers.nexon.com/apiReference/Misc/List-1) GetAllTilePositions() |
| --- |
| 현재 타일맵에 존재하는 모든 타일들의 위치 리스트를 반환합니다. |

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetTile(int32 cellPositionX, int32 cellPositionY) |
| --- |
| 해당 위치에 있는 타일의 정보를 반환합니다. |

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetTile([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| 해당 위치에 있는 타일의 정보를 반환합니다. |

| void RemoveTile(int32 cellPositionX, int32 cellPositionY) |
| --- |
| 해당 위치의 타일을 제거합니다. |

| void RemoveTile([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| 해당 위치의 타일을 제거합니다. |

| void Reset() |
| --- |
| 타일맵을 초기 상태로 되돌립니다. |

| void SetTile(int32 tileIndex, int32 cellPositionX, int32 cellPositionY) |
| --- |
| 해당 위치에 타일을 배치합니다. tileIndex는 타일의 인덱스 번호입니다. |

| void SetTile(int32 tileIndex, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| 해당 위치에 타일을 배치합니다. tileIndex는 타일의 인덱스 번호입니다. |

| void SetTile([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, int32 cellPositionX, int32 cellPositionY) |
| --- |
| 해당 위치에 타일을 배치합니다. tileName은 타일의 Name 프로퍼티입니다. |

| void SetTile([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| 해당 위치에 타일을 배치합니다. tileName은 타일의 Name 프로퍼티입니다. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) ToCellPosition(float worldPositionX, float worldPositionY) |
| --- |
| 실수형 월드 공간 좌표를 정수형 타일맵 공간 좌표로 변환합니다. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) ToCellPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPosition) |
| --- |
| 실수형 월드 공간 좌표를 정수형 타일맵 공간 좌표로 변환합니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPosition(int32 cellPositionX, int32 cellPositionY) |
| --- |
| 정수형 타일맵 공간 좌표를 실수형 월드 공간 좌표로 변환합니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPosition([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| 정수형 타일맵 공간 좌표를 실수형 월드 공간 좌표로 변환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

밟고 있는 타일에 따라 이동 속력이 바뀌는 예제입니다.

```
Method:
[client only]
void OnUpdate ()
{
	if _UserService.LocalPlayer ~= self.Entity then
		return
	end
	 
	local transform = self.Entity.TransformComponent
	local movement = self.Entity.MovementComponent
	local tilemap = self.Entity.CurrentMap:GetFirstChildComponentByTypeName("RectTileMapComponent")
	 
	local worldPos = transform.WorldPosition
	 
	-- 월드 공간 좌표를 타일맵 공간 좌표로 변환
	local cellPos = tilemap:ToCellPosition(worldPos)
	  
	-- 현재 위치에 있는 타일 정보 조회
	local tileInfo = tilemap:GetTile(cellPos)
	 
	if tileInfo == nil then
		return
	end
	 
	if (tileInfo.Name == "Fast Tile") then
		movement.InputSpeed = 2.4
	elseif (tileInfo.Name == "Slow Tile") then
		movement.InputSpeed = 1.2
	end
}
```

# SeeAlso

- [MovementComponent](https://mod-developers.nexon.com/apiReference/Components/MovementComponent)
- [TransformComponent](https://mod-developers.nexon.com/apiReference/Components/TransformComponent)
- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)
- [RectTileMap의 활용](/docs?postId=589)
- [SideViewRectTile모드로 맵 만들기](/docs?postId=758)

Update 2025-10-28 PM 02:21


# RevoluteJointComponent

RevoluteJoint를 생성, 삭제합니다. 연결된 강체끼리의 상대적 회전을 제어합니다.

# Properties

| [SyncList<RevoluteJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Joint 정보를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Joint를 추가합니다. 성공 시 index, 실패 시 -1을 반환합니다 |

| void DestroyJoint(int32 index) |
| --- |
| 순번이 index에 해당하는 Joint를 제거합니다. |

| int32 GetJointsCount() |
| --- |
| joint 수를 반환합니다. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| 순번이 index에 해당하는 Joint의 CollideConnected 값을 설정합니다. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorA 값을 설정합니다. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| 순번이 index에 해당하는 Joint의 LocalAnchorB 값을 설정합니다. |

| void SetLowerAngle(int32 index, float lowerAngle) |
| --- |
| 순번이 index에 해당하는 Joint의 LowerAngle 값을 설정합니다. |

| void SetMaxMotorTorque(int32 index, float maxMotorTorque) |
| --- |
| 순번이 index에 해당하는 Joint의 MaxMotorTorque 값을 설정합니다. |

| void SetMotorEnable(int32 index, boolean enable) |
| --- |
| 순번이 index에 해당하는 Joint의 MotorEnable 값을 설정합니다. |

| void SetMotorSpeed(int32 index, float speed) |
| --- |
| 순번이 index에 해당하는 Joint의 MotorSpeed 값을 설정합니다. |

| void SetUpperAngle(int32 index, float upperAngle) |
| --- |
| 순번이 index에 해당하는 Joint의 UpperAngle 값을 설정합니다. |

| void SetUseLimits(int32 index, boolean useLimits) |
| --- |
| 순번이 index에 해당하는 Joint의 UseLimits 값을 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

#### SetMotorSpeed

W, S를 누르면 TargetEntityRef로 지정한 Entity를 회전시키는 예제입니다.

```
Method:
[server]
void GoUp ( )
{
	self.Entity.RevoluteJointComponent:SetMotorSpeed(1, 10)
}

[server]
void GoDown ( )
{
	self.Entity.RevoluteJointComponent:SetMotorSpeed(1, -10)
}

[server]
void Stop ( )
{
	self.Entity.RevoluteJointComponent:SetMotorSpeed(1, 0)
}


Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.W then
		self:GoUp()
	elseif key == KeyboardKey.S then
		self:GoDown()
	end
}

[service: InputService]
HandleKeyUpEvent (KeyUpEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	self:Stop()
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)
- [KeyUpEvent](https://mod-developers.nexon.com/apiReference/Events/KeyUpEvent)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [다양한 물리 joint 활용하기](/docs?postId=760)
- [물리 사용하기](/docs?postId=757)
- [엔티티에 물리 적용하기](/docs?postId=761)

Update 2025-10-28 PM 02:21


# RigidbodyComponent

메이플스토리 움직임을 적용합니다. 중력 및 가·감속 영향을 받습니다.

# Properties

| float AirAccelerationX |
| --- |
| 공중에서의 속도를 보정합니다. 값이 증가할 수록 공중에서의 이동 속도가 빨라집니다. |

| float AirDecelerationX |
| --- |
| 공중에서의 입력이 없을 경우 X축의 이동 속도가 얼마나 빠르게 멈추는지를 보정합니다. |

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true인 경우 회전하거나 기울어진 사다리를 탄 캐릭터는 사다리의 모습을 따릅니다. false인 경우 캐릭터는 사다리의 기울기, 회전에 영향을 받지 않습니다. |

| float DownJumpSpeed |
| --- |
| 아래 점프 시 위로 튀어 오르는 속도를 조절합니다. |

| boolean EnableKinematicMoveJump |
| --- |
| KinematicMove가 true일 경우 점프를 허용하거나 허용하지 않습니다. |

| float FallSpeedMaxX |
| --- |
| 공중에서의 X축 최대 속도 제한값을 보정합니다. |

| float FallSpeedMaxY |
| --- |
| 공중에서의 Y축 최대 속도 제한값을 보정합니다. |

| float Gravity |
| --- |
| 중력값입니다. 공중에서 이동할 때에 지상으로 얼마나 빨리 떨어질 것인가와 연관이 있습니다. 입력값이 클수록 더욱 더 빠르게 떨어집니다. |

| boolean IgnoreMoveBoundary |
| --- |
| true인 경우 지형에 의해 생성된 맵 영역을 벗어날 수 있습니다. |

| boolean IsBlockVerticalLine |
| --- |
| true일 경우 기본 이동과 다르게 세로 지형이 무조건 막힙니다. 벽 같은 지형을 무조건 통과할 수 없도록 할 때 사용합니다. |

| boolean IsolatedMove |
| --- |
| 값이 true일 경우 발판 끝에 도달해도 떨어지지 않습니다. 점프 같은 외부 이동으로는 발판을 벗어날 수 있습니다. |

| float JumpBias |
| --- |
| 캐릭터가 얼마만큼 공중에 처음 뜰지를 설정합니다. |

| boolean KinematicMove |
| --- |
| true인 경우 탑다운 방식의 상하좌우 이동으로 변경됩니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) KinematicMoveAcceleration |
| --- |
| 이동 속력을 설정합니다. KinematicMove가 true일 때 동작합니다. |

| [AutomaticLayerOption](https://mod-developers.nexon.com/apiReference/Enums/AutomaticLayerOption) LayerSettingType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| RigidbodyComponent와 foothold, 사다리, 로프의 SortingLayer 값의 관계를 설정합니다. |

| float Mass |
| --- |
| 질량을 설정합니다. 값이 클 수록 가감속이 느려지고 외부 요인에 대한 반응성이 낮아집니다. 0보다 큰 값만 설정할 수 있습니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 이동에 필요한 입력값입니다. 주로 MovementComponent에서 입력을 제어합니다. X가 양수일 경우 오른쪽, Y가 양수일 경우 위쪽을 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RealMoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 직전에 이동된 양을 나타냅니다. 읽기 전용이며 이동량 정보가 필요할 때 사용합니다. 이동 주체에 따라 유효한 실행공간이 존재하며 LocalPlayer는 Client, 그 외는 Server에서 값을 가집니다. |

| float WalkAcceleration |
| --- |
| 지형 이동 시 가감속 값을 나타냅니다. 입력값이 클수록 더욱 더 빠르게 최대 속도에 도달합니다. |

| float WalkDrag |
| --- |
| 지형 이동 시 미끄러짐에 저항하는 힘입니다. 입력값이 클수록 미끄러지지 않고 빠르게 멈추게 됩니다. 맵, 지형, 캐릭터의 속성을 계산해 최종 적용 값은 0.5부터 2 사이의 범위입니다. |

| float WalkJump |
| --- |
| 점프 시 얼마나 높게 뛰어오를지에 대한 값입니다. 값이 클수록 높게 뜁니다. |

| float WalkSlant |
| --- |
| 지형 이동 시 경사를 얼마나 잘 넘을 수 있는지와 연관된 값입니다. 입력값이 크면 급경사를 넘어갈 수 있습니다. 0부터 1 사이의 값을 가집니다. |

| float WalkSpeed |
| --- |
| 지형 이동 시 최대 이동 속도 값을 조절합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void AddForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| 엔티티에 가해지는 힘을 더합니다. 엔티티는 기존 힘에 추가로 더해진 힘의 방향으로 가·감속 운동을 하게 됩니다. |

| void AttachTo([string](https://mod-developers.nexon.com/apiReference/Lua/string) entityId, [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) offset) |
| --- |
| 이 엔티티를 entityId에 해당하는 엔티티에게 붙힙니다. 이로서 이 엔티티는 물리 동작을 하지 않고 붙혀진 엔티티의 이동에 종속되게 됩니다. |

| void Detach() |
| --- |
| RigidbodyComponent:AttachTo(string, Vector3)로 다른 엔티티에 붙어 있던 엔티티를 떼어냅니다. |

| boolean DownJump() |
| --- |
| 아래 점프를 수행합니다. 아래 점프는 지형 위에 있을 때만 유효합니다. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetCurrentFoothold() |
| --- |
| 현재 밟고 있는 Foothold를 반환합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetCurrentFootholdPerpendicular() |
| --- |
| 밟고 있는 지형의 수직선을 반환합니다. |

| boolean IsOnGround() |
| --- |
| 지형 위에 서 있는지를 확인합니다. |

| boolean JustJump([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) jumpRate) |
| --- |
| 대상을 점프시킵니다. |

| void PositionReset() |
| --- |
| 누적된 위치 계산 정보를 삭제하고, 현재 위치를 기반으로 새롭게 계산합니다. |

| boolean PredictFootholdEnd(float distance, boolean isFoward) |
| --- |
| 지금 밟고 있는 발판에서 distance만큼 이동할 수 있는지 확인합니다. isForward가 true면 오른쪽, false면 왼쪽 방향을 확인합니다.<br>현재 위치에서 발판 끝까지 거리가 distance보다 멀면 true, 가까우면 false를 반환합니다. 발판을 밟고 있지 않다면 false를 반환합니다. |

| void SetForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| 엔티티에 가해지는 힘을 설정합니다. 엔티티는 설정한 힘의 방향으로 가·감속 운동을 하게 됩니다. |

| void SetForceReserve([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| 힘을 즉시 가하지 않고, 현재 프레임에서의 이동이 끝난 후 주어진 입력값으로 힘을 대체합니다. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 로컬 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void SetUseCustomMove(boolean isUse) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. RigidbodyComponent의 Enable 프로퍼티를 사용하세요. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 월드 좌표 기준으로 엔티티의 위치를 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [FootholdCollisionEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdCollisionEvent) |
| --- |
| RigidBodyComponent가 발판에 충돌했을 때 발생하는 이벤트입니다. |

| [FootholdEnterEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdEnterEvent) |
| --- |
| 엔티티가 Foothold에 붙었을 때 발생하는 이벤트입니다. |

| [FootholdLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdLeaveEvent) |
| --- |
| 엔티티가 Foothold에서 떨어졌을 때 발생하는 이벤트입니다. |

| [RigidbodyAttachEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyAttachEvent) |
| --- |
| 엔티티가 RigidbodyComponent:AttachTo(string, Vector3)를 통해 특정 엔티티에 붙었을 때 발생하는 이벤트입니다. 플레이어는 클라이언트에서 발생하고, 그 외의 엔티티는 서버 공간에서 발생합니다. |

| [RigidbodyClimbableAttachStartEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyClimbableAttachStartEvent) |
| --- |
| Avatar가 사다리, 로프를 타기 전에 발생하는 이벤트입니다. |

| [RigidbodyClimbableDetachEndEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyClimbableDetachEndEvent) |
| --- |
| Avatar가 사다리, 로프에서 떨어진 후에 발생하는 이벤트입니다. |

| [RigidbodyDetachEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyDetachEvent) |
| --- |
| 엔티티가 RigidbodyComponent:Detach()를 통해 Attach 상태가 해제될 때 발생하는 이벤트입니다. 플레이어는 클라이언트에서 발생하고, 그 외의 엔티티는 서버 공간에서 발생합니다. |

| [RigidbodyKinematicMoveJumpEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyKinematicMoveJumpEvent) |
| --- |
| KinematicMove 프로퍼티가 true인 경우 점프하거나 착지하면 발생하는 이벤트입니다. |

# Examples

#### AttachTo

특정 오브젝트에 닿으면 Attach 되고 attach 된 상태에서 3초 후 Dtach됩니다.

```
Property:
[Sync]
number time = 0
[Sync]
boolean isAttached = false

Method:
[client]
void AttachTo (string entityId)
{
	self.Entity.RigidbodyComponent:AttachTo(entityId, Vector3.zero)
	self.isAttached = true
}

[client only]
void OnUpdate (number delta)
{
	if self.isAttached == false then
		return
	end
	 
	self.time = self.time + delta
	 
	if self.time >= 3.0 then
		self.Entity.RigidbodyComponent:Detach()
		self.time = 0
		self.isAttached = false
	end
}

Event Handler:
[self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TriggerComponent
	-- Space: Server, Client
	---------------------------------------------------------
	
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	---------------------------------------------------------
	if TriggerBodyEntity.Name == "Name" then
		self:AttachTo(TriggerBodyEntity.Id)
	end
}
```

#### PredictFootholdEnd

발판의 끝점이 특정 distance만큼 떨어져 있는지 판별합니다. 캐릭터가 밟고 있는 발판에서 양의 방향으로 발판의 끝이 10보다 가까우면 true, 멀다면 false를 반환합니다. 이에 따라 Entity의 Enable이 변경됩니다.

```
Method:
[client only]
void OnUpdate (number delta)
{
	local entity = _EntityService:GetEntityByPath(EntityPath)
	if self.Entity.RigidbodyComponent:PredictFootholdEnd(10, true) then
		entity.Enable = true
	else
		entity.Enable = false
	end
}
```

# SeeAlso

- [TriggerComponent](https://mod-developers.nexon.com/apiReference/Components/TriggerComponent)
- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [메이플 이동 개념 이해하기](/docs?postId=750)

Update 2025-12-02 PM 01:55


# ScrollLayoutGroupComponent

스크롤뷰와 관련된 컴포넌트들을 하나로 제어할 수 있습니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CellSize |
| --- |
| 자식 UI Entity의 고정 크기입니다. Grid 타입 전용입니다. |

| [ChildAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/ChildAlignmentType) ChildAlignment |
| --- |
| Entity에 여유 공간이 있을 때 자식 UI Entity의 정렬 방식을 설정합니다. Vertical, Grid 타입 전용입니다. |

| [GridLayoutConstraint](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutConstraint) Constraint |
| --- |
| 행, 열 개수에 대한 제약 사항입니다. Grid 타입 전용입니다. |

| int32 ConstraintCount |
| --- |
| 제약 사항에 따라 고정하고자 하는 행 또는 열의 개수입니다. Grid 타입 전용입니다. |

| [ChildAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/ChildAlignmentType) GridChildAlignment |
| --- |
| Entity에 여유 공간이 있을 때 자식 UI Entity의 정렬 방식을 설정합니다. Grid 타입 전용입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GridSpacing |
| --- |
| 자식 UI Entity간의 가로, 세로 간격입니다. Grid 타입 전용입니다. |

| [HorizontalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/HorizontalScrollBarDirection) HorizontalScrollBarDirection |
| --- |
| 가로 스크롤바의 방향입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되어 있는지 여부를 나타냅니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| 레이아웃 그룹의 상하좌우의 여유 공간을 설정합니다. |

| boolean ReverseArrangement |
| --- |
| 기존 정렬과 반대로 정렬할지 여부입니다. Vertical, Horizontal 타입 전용입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarBackgroundColor |
| --- |
| 스크롤바의 배경 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarBgImageRUID |
| --- |
| 스크롤바의 배경 이미지입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarHandleColor |
| --- |
| 스크롤바의 핸들 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarHandleImageRUID |
| --- |
| 스크롤바의 핸들 이미지입니다. |

| float ScrollBarThickness |
| --- |
| 스크롤바 영역의 두께입니다. |

| [ScrollBarVisibility](https://mod-developers.nexon.com/apiReference/Enums/ScrollBarVisibility) ScrollBarVisible |
| --- |
| 스크롤바의 표시 여부를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| float Spacing |
| --- |
| 자식 UI Entity간의 간격입니다. Vertical, Horizontal 타입 전용입니다. |

| [GridLayoutAxis](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutAxis) StartAxis |
| --- |
| 자식 UI Entity의 정렬 방향입니다. Grid 타입 전용입니다. |

| [GridLayoutCorner](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutCorner) StartCorner |
| --- |
| 자식 UI Entity의 정렬 시작 위치입니다. Grid 타입 전용입니다. |

| [LayoutGroupType](https://mod-developers.nexon.com/apiReference/Enums/LayoutGroupType) Type |
| --- |
| 레이아웃 그룹의 정렬 형식을 설정합니다. |

| boolean UseScroll |
| --- |
| 스크롤 기능 사용 여부를 설정합니다. |

| [VerticalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/VerticalScrollBarDirection) VerticalScrollBarDirection |
| --- |
| 세로 스크롤바의 방향입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetScrollNormalizedPosition() |
| --- |
| 스크롤바의 정규화된 위치를 반환합니다. |

| float GetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| 지정한 방향 스크롤바의 정규화된 위치를 반환합니다. |

| void ResetScrollPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| 지정한 축의 스크롤바 위치를 처음 위치로 이동합니다. 스크롤바의 방향에 따라 처음 위치가 다릅니다. |

| void SetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis, float value) |
| --- |
| 지정한 축의 스크롤바 위치를 지정한 정규화된 위치로 이동합니다. 스크롤바의 방향에 따라 0, 1의 방향이 다릅니다. |

| void SetScrollPositionByItemIndex(int32 index) |
| --- |
| 스크롤바 위치를 특정 index의 자식 UI Entity가 보이는 위치로 이동합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [ScrollPositionChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ScrollPositionChangedEvent) |
| --- |
| 스크롤이 가능한 UI Entity에서 스크롤 위치가 변경될 때 발생하는 이벤트입니다. UI Entity에 ScrollLayoutGroupComponent 또는 GridViewComponent가 있어야 이벤트가 발생합니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

`ScrollLayoutGroupComponent`의 `CellSize`를 변경하고, 1열로 고정한 다음, `ItemEntity`를 `InitItemCount` 만큼 복제해서 `ScrollLayoutGroupComponent`에 추가하는 예제입니다.

이 Component를 `ScrollLayoutGroupComponent`가 있는 Entity에 추가하고, `ItemEntity`, `InitItemCount` 프로퍼티 값을 설정한 후 테스트 할 수 있습니다.

```
Property:
[None]
Entity ItemEntity = "nil"
[None]
number InitItemCount = 0
  
Method:
[ClientOnly]
void OnBeginPlay()
{
	self.ItemEntity = self.Entity:GetChildByName("Scroll_Image")
	self.InitItemCount = 10
	 
	self:Init()
}
  
void Init()
{
	if not self.ItemEntity then
		return
	end
	
	local scrollLayoutComp = self.Entity.ScrollLayoutGroupComponent
	if scrollLayoutComp then
		scrollLayoutComp.CellSize = Vector2(200, 200)
		scrollLayoutComp.Constraint = GridLayoutConstraint.FixedColumnCount
		scrollLayoutComp.ConstraintCount = 1
	end
	 
	local i = 0
	for i = 0, self.InitItemCount, 1 do
		local itemslot = _SpawnService:Clone("Item", self.ItemEntity, self.Entity)
	end
}
```

# SeeAlso

- [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity)
- [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2)
- [SpawnService](https://mod-developers.nexon.com/apiReference/Services/SpawnService)

Update 2025-12-02 PM 01:55


# SideviewbodyComponent

횡스크롤 방식의 이동 및 점프, 렉트 타일과의 충돌 기능을 제공합니다. 중력의 영향을 받으며, 가·감속 영향은 받지 않습니다. 현재 타일맵 모드가 SideViewRectTileMap일 때 동작합니다.

# Properties

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true인 경우 회전하거나 기울어진 사다리를 탄 캐릭터는 사다리의 모습을 따릅니다. false인 경우 캐릭터는 사다리의 기울기, 회전에 영향을 받지 않습니다. |

| float DownJumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 아래 점프 시 위로 튀어 오르는 속력을 조절합니다. 값이 클수록 더 높게 점프합니다. |

| boolean EnableDownJump ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 아래 점프 기능을 키거나 끕니다. |

| float JumpDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 속력 감소량을 조절합니다. 값이 클수록 지면에 더 빨리 떨어집니다. |

| float JumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 점프 시 튀어 오르는 속력을 조절합니다. 값이 클수록 더 높게 점프합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 이동 속도를 설정합니다. MovementComponent가 이동 제어를 위해 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetUnderfootTile() |
| --- |
| 현재 밟고 있는 타일 정보를 확인하고, 밟은 타일이 없을 경우 nil을 반환합니다. 수평으로 나란히 배치된 두 타일의 경계 위에 있는 경우 왼쪽 타일을 반환합니다. |

| boolean IsOnGround() |
| --- |
| 현재 지면에 닿아 있는 상태인지 아닌지 확인합니다. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 로컬 좌표 기준으로 엔티티의 위치를 설정합니다. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| 월드 좌표 기준으로 엔티티의 위치를 설정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [RectTileCollisionBeginEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionBeginEvent) |
| --- |
| 충돌 가능한 타일과 접촉했을 때 발생하는 이벤트입니다. |

| [RectTileCollisionEndEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionEndEvent) |
| --- |
| 충돌한 타일에서 벗어날 때 발생하는 이벤트입니다. |

| [RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) |
| --- |
| 특정 사각형 타일에 진입했을 때 발생하는 이벤트입니다. |

| [RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) |
| --- |
| 특정 사각형 타일에서 벗어났을 때 발생하는 이벤트입니다. |

# Examples

#### GetWallTile

측면에 닿은 벽 타일을 감지하는 예제입니다.

```
Property:
[None]
table WallTile = {}
 
Method:
[client only]
table GetWallTile ()
{
	return self.WallTile
}

Event Handler:
[client only] [self]
HandleRectTileCollisionBeginEvent (RectTileCollisionBeginEvent event)
{
	----------------- Native Emitter Info ------------------
	-- Emitter: KinematicbodyComponent
	-- Space: Server, Client
	--------------------------------------------------------
	-- Emitter: SideviewbodyComponent
	-- Space: Server, Client
	--------------------------------------------------------
	      
	-- Parameters
	local Entity = event.Entity
	local Normal = event.Normal
	local Point = event.Point
	local TileInfo = event.TileInfo
	local TileMap = event.TileMap
	local TilePosition = event.TilePosition
	--------------------------------------------------------
	    
	if Normal == Vector2.left or Normal == Vector2.right then
		self.WallTile ={
			info = TileInfo,
			position = TilePosition:Clone(),
			normal = Normal:Clone(),
			tilemap = TileMap
		}
	end 
}
  
[client only] 
[self]
HandleRectTileCollisionEndEvent ( RectTileCollisionEndEvent event )
{
	----------------- Native Emitter Info ------------------
	-- Emitter: KinematicbodyComponent
	-- Space: Server, Client
	--------------------------------------------------------
	-- Emitter: SideviewbodyComponent
	-- Space: Server, Client
	--------------------------------------------------------
	      
	-- Parameters
	local Entity = event.Entity
	local Normal = event.Normal
	local Point = event.Point
	local TileInfo = event.TileInfo
	local TileMap = event.TileMap
	local TilePosition = event.TilePosition
	--------------------------------------------------------
	      
	local currWallTile = self.WallTile
	      
	if currWallTile ~= nil and currWallTile.position == TilePosition then
		self.WallTile = nil
	end
}
```

# SeeAlso

- [RectTileCollisionBeginEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionBeginEvent)
- [RectTileCollisionEndEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionEndEvent)
- [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2)
- [SideViewRectTile모드로 맵 만들기](/docs?postId=758)
- [SideViewRectTileMap에서 캐릭터 이동 제어](/docs?postId=759)

Update 2025-10-28 PM 02:21


# SkeletonGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Spine으로 제작된 스켈레톤 리소스를 UI에 그리고 제어할 수 있는 기능을 제공합니다. Spine 4.1로 제작된 리소스만 사용할 수 있습니다.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) AnimationNames |
| --- |
| 1번 트랙에 애니메이션을 설정합니다. 애니메이션은 인덱스 순서대로 재생됩니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 색상을 설정합니다. |

| boolean FlipX |
| --- |
| X축 반전 여부를 설정합니다. |

| boolean FlipY |
| --- |
| Y축 반전 여부를 설정합니다. |

| boolean Loop |
| --- |
| 애니메이션 반복 재생 여부를 설정합니다. |

| float PlayRate |
| --- |
| 애니메이션 재생 속도를 설정합니다. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveMode |
| --- |
| 스켈레톤 리소스의 비율, 피봇, 크기 등이 유지되는 방식을 설정합니다. |

| boolean RaycastTarget |
| --- |
| true로 설정할 경우 화면 터치 또는 마우스 클릭 대상이 되며, 뒤에 가려진 UI는 화면 터치와 마우스 클릭 입력을 받지 못합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SkeletonRUID |
| --- |
| 스켈레톤 리소스의 RUID를 설정합니다. |

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) SkinNames |
| --- |
| 스킨을 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void AddAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 트랙에 애니메이션을 추가합니다. animationClip에는 트랙 번호, 애니메이션 이름, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

| void AddEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 트랙에 빈 애니메이션을 추가합니다. animationClip에는 트랙 번호, 기타 속성을 설정합니다. 1번 트랙은 사용할 수 없습니다. |

| void ClearTrack(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 트랙을 비웁니다. 1번 트랙은 사용할 수 없습니다. |

| [SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) GetCurrentAnimation(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 현재 재생 중인 애니메이션을 반환합니다. |

| void SetAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 트랙을 비우고 애니메이션을 추가합니다. animationClip에는 트랙 번호, 애니메이션 이름, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

| void SetAttachment([string](https://mod-developers.nexon.com/apiReference/Lua/string) slotName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attachmentName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 슬롯에 어태치먼트를 붙입니다. attachmentName이 nil이면 슬롯에서 어태치먼트를 제거합니다. |

| void SetEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 트랙을 비우고 빈 애니메이션을 추가합니다. animationClip에는 트랙 번호, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [SkeletonAnimationCompleteEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationCompleteEvent) |
| --- |
| 스켈레톤 애니메이션의 재생이 완료되었을 때 발생하는 이벤트입니다. 반복 재생되는 경우, 재생이 완료될 때마다 발생합니다. |

| [SkeletonAnimationEndEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationEndEvent) |
| --- |
| 스켈레톤 애니메이션이 전환될 때, 종료되는 애니메이션에서 발생하는 이벤트입니다. |

| [SkeletonAnimationStartEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationStartEvent) |
| --- |
| 스켈레톤 애니메이션이 전환될 때, 새롭게 재생되는 애니메이션에서 발생하는 이벤트입니다. |

| [SkeletonAnimationTimelineEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationTimelineEvent) |
| --- |
| 스켈레톤 애니메이션이 재생 중일 때, 애니메이션 타임라인에 등록된 이벤트가 감지되면 발생하는 이벤트입니다. |

Update 2025-12-03 PM 05:12


# SkeletonRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Spine으로 제작된 스켈레톤 리소스를 그리고 제어할 수 있는 기능을 제공합니다. Spine 4.1로 제작된 리소스만 사용할 수 있습니다.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) AnimationNames ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 1번 트랙에 애니메이션을 설정합니다. 애니메이션은 인덱스 순서대로 재생됩니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 색상을 설정합니다. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Y축을 기준으로 반전 여부를 결정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 맵 레이어 이름을 지정했을 때 자동으로 치환되지 않습니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 애니메이션 반복 재생 여부를 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 애니메이션 재생 속도를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SkeletonRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스켈레톤 리소스의 RUID를 설정합니다. |

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) SkinNames ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스킨을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void AddAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| 트랙에 애니메이션을 추가합니다. animationClip에는 트랙 번호, 애니메이션 이름, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

| void AddEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| 트랙에 빈 애니메이션을 추가합니다. animationClip에는 트랙 번호, 기타 속성을 설정합니다. 1번 트랙은 사용할 수 없습니다. |

| void ClearTrack(int32 trackIndex) |
| --- |
| 트랙을 비웁니다. 1번 트랙은 사용할 수 없습니다. |

| [SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) GetCurrentAnimation(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 현재 재생 중인 애니메이션을 반환합니다. |

| void SetAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| 트랙을 비우고 애니메이션을 추가합니다. animationClip에는 트랙 번호, 애니메이션 이름, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

| void SetAttachment([string](https://mod-developers.nexon.com/apiReference/Lua/string) slotName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attachmentName) |
| --- |
| 슬롯에 어태치먼트를 붙입니다. attachmentName이 nil이면 슬롯에서 어태치먼트를 제거합니다. |

| void SetEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| 트랙을 비우고 빈 애니메이션을 추가합니다. animationClip에는 트랙 번호, 기타 속성을 지정합니다. 1번 트랙은 사용할 수 없습니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SkeletonAnimationCompleteEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationCompleteEvent) |
| --- |
| 스켈레톤 애니메이션의 재생이 완료되었을 때 발생하는 이벤트입니다. 반복 재생되는 경우, 재생이 완료될 때마다 발생합니다. |

| [SkeletonAnimationEndEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationEndEvent) |
| --- |
| 스켈레톤 애니메이션이 전환될 때, 종료되는 애니메이션에서 발생하는 이벤트입니다. |

| [SkeletonAnimationStartEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationStartEvent) |
| --- |
| 스켈레톤 애니메이션이 전환될 때, 새롭게 재생되는 애니메이션에서 발생하는 이벤트입니다. |

| [SkeletonAnimationTimelineEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationTimelineEvent) |
| --- |
| 스켈레톤 애니메이션이 재생 중일 때, 애니메이션 타임라인에 등록된 이벤트가 감지되면 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-12-02 PM 01:55


# SliderComponent

최소, 최대 범위 내에서 값을 설정하고, 해당 값을 그래픽으로 나타냅니다.

# Properties

| [SliderDirection](https://mod-developers.nexon.com/apiReference/Enums/SliderDirection) Direction |
| --- |
| 최솟값에서 최댓값이 그래픽으로 나타낼 방향을 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FillRectColor |
| --- |
| Value를 그래픽으로 나타낼 영역의 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) FillRectImageRUID |
| --- |
| Value를 그래픽으로 나타낼 영역의 이미지 RUID입니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) FillRectPadding |
| --- |
| Value를 그래픽으로 나타낼 영역의 상하좌우 여유 공간을 설정합니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) HandleAreaPadding |
| --- |
| 슬라이더 핸들로 이동 가능한 영역의 상하좌우 여유 공간입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) HandleColor |
| --- |
| 슬라이더 핸들의 색상입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) HandleImageRUID |
| --- |
| 슬라이더 핸들의 이미지 RUID입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) HandleSize |
| --- |
| 슬라이더 핸들의 크기입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되어 있는지 여부를 나타냅니다. |

| float MaxValue |
| --- |
| Value의 최댓값입니다. |

| float MinValue |
| --- |
| Value의 최솟값입니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| boolean UseHandle |
| --- |
| 핸들의 사용 여부를 설정합니다. |

| boolean UseIntegerValue |
| --- |
| Value를 정수로만 사용할 건지 여부를 설정합니다. |

| float Value |
| --- |
| 현재값입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SliderValueChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SliderValueChangedEvent) |
| --- |
| Slider 값이 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

`SliderComponent`의 값 범위를 0 ~ 100으로 설정하고, 값을 `TextComponent`에 표시하는 예제입니다.

```
Property:
[None]
Entity TextEntity = EntityPath
  
Method:
[client only]
void OnBeginPlay()
{
	local sliderComp = self.Entity.SliderComponent
	if not sliderComp then
		return
	end
	 
	sliderComp.UseIntegerValue = true
	sliderComp.MaxValue = 100
	sliderComp.MinValue = 0
	sliderComp.Value = 0
	 
	self:SetSliderText(sliderComp.Value)
}
  
void SetSliderText(number sliderValue)
{
	if not self.TextEntity then
		return
	end
  
	local textComp = self.TextEntity.TextComponent
	if not textComp then
		return
	end
  
	textComp.Text = string.format("%d", sliderValue)
}
  
Event Handler:
[self]
HandleSliderValueChangedEvent (SliderValueChangedEvent event)
{
	-- Parameters
	local Value = event.Value
	--------------------------------------------------------
	self:SetSliderText(sliderValue)
}
```

# SeeAlso

- [TextComponent](https://mod-developers.nexon.com/apiReference/Components/TextComponent)
- [SliderValueChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SliderValueChangedEvent)

Update 2025-08-27 PM 04:56


# SoundComponent

효과음 또는 배경음악을 재생하고 관리합니다.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) AudioClipRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 음원 재생에 사용할 AudioClipRUID를 설정합니다. |

| boolean Bgm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 배경음악으로 재생될지 여부를 설정합니다. |

| float HearingDistance ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 소리가 들리는 리스너 엔티티와의 최대 거리를 설정합니다. |

| boolean KeepBGM ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이전 BGM과 현재 BGM이 같다면 음원을 이어서 재생합니다. Bgm, PlayOnEnable 프로퍼티가 true이고 SoundComponent가 활성화될 때 적용됩니다.<br>SoundService의 PlayBGM() 함수들로 호출할 때는 BGM이 이어서 재생되지 않습니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 반복 재생 여부를 설정합니다. |

| boolean Mute ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 음소거 상태를 설정합니다. |

| float Pitch ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 음원의 음높이와 재생 속도를 설정합니다. 값이 커질수록 음높이가 높아지며, 재생 속도가 빨라집니다. 0 이상, 3 이하의 값만 설정할 수 있습니다. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Enable이 활성화될 때 음원 재생 여부를 설정합니다. |

| boolean SetCameraAsListener ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 화면의 중앙과 음원 사이의 거리에 따라 소리 크기를 조절합니다. SetListenerEntity 함수를 통해 따로 지정한 리스너 엔티티가 있을 때는 이 프로퍼티의 활성 여부와 관계 없이 해당 엔티티가 리스너가 됩니다. |

| float Volume ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 음량을 설정합니다. 0 이상, 1 이하의 값만 설정할 수 있습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| float GetTimePosition() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 음원의 현재 재생 위치를 초 단위로 반환합니다. 오디오 클립이 로드되지 않은 경우 -1을 반환합니다. |

| float GetTotalTime() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 음원의 전체 길이를 초 단위로 반환합니다. 오디오 클립이 로드되지 있지 않은 경우 -1을 반환합니다. |

| boolean IsAudioClipLoaded() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| AudioClipRUID에 설정한 오디오 클립이 로드되어 있는지 여부를 반환합니다. GetTimePosition(), GetTotalTime(), SetTimePosition(timeInSecond, targetUserId) 함수는 오디오 클립이 로드되어 있지 않은 경우 실행할 수 없습니다. |

| boolean IsPlaying([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 음원이 재생 중인지 여부를 확인합니다. |

| boolean IsSyncedPlaying() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 동기화되는 음원이 재생 중인지 여부를 확인합니다. PlaySyncedSound가 호출된 이후, 따로 StopSyncedSound를 호출한 경우나 음원이 끝까지 재생되어 정지된 경우가 아니라면 true를 반환합니다. |

| void Pause([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 음원 재생을 일시 정지합니다. |

| void Play([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 음원을 재생합니다. |

| void PlaySyncedSound() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| 동기화되는 음원을 재생합니다. BGM 프로퍼티가 true인 경우 작동하지 않습니다. |

| void Resume([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 음원 재생을 재개합니다. |

| void SetListenerEntity([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) entity, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 리스너 엔티티를 설정합니다. 리스너 엔티티와의 거리가 멀어질 수록 음량이 작아집니다. |

| void SetTimePosition(float timeInSecond, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 재생 위치를 변경합니다. 초 단위를 사용합니다. 오디오 클립이 로드되지 않은 경우 동작하지 않습니다. |

| void Stop([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| 음원 재생을 정지합니다. |

| void StopSyncedSound() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| 동기화되는 음원 재생을 정지합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [SoundPlayStateChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SoundPlayStateChangedEvent) |
| --- |
| SoundService의 BGM, SoundComponent의 효과음 재생 상태가 변경될 때 발생하는 이벤트입니다. |

# Examples

음원 리소스(AudioClipRUID)를 설정하고 로컬 플레이어를 리스너로 선택합니다. SoundComponent 위치로 다가갈수록 설정한 소리가 커지는 환경을 구성했습니다.

```
Method:
[client only]
void OnBeginPlay ()
{
	self.Entity.SoundComponent.AudioClipRUID = "000000" 
	self.Entity.SoundComponent.Loop = true
	self.Entity.SoundComponent:SetListenerEntity(_UserService.LocalPlayer)
	self.Entity.SoundComponent:Play()
}
```

# SeeAlso

- [UserService](https://mod-developers.nexon.com/apiReference/Services/UserService)
- [배경음악 변경하기](/docs?postId=117)
- [효과음 만들기](/docs?postId=578)

Update 2025-10-28 PM 02:21


# SpawnLocationComponent

SpawnLocation Model에서만 사용하는 특수 컴포넌트입니다.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# SpriteGUIRendererComponent

UI에 스프라이트 또는 애니메이션 클립을 출력합니다.

# Properties

| [SpriteAnimClipPlayType](https://mod-developers.nexon.com/apiReference/Enums/SpriteAnimClipPlayType) AnimClipPlayType |
| --- |
| 애니메이션 클립의 재생 방식을 설정합니다. 한 번만 재생, 반복 재생 등을 설정할 수 있습니다. |

| int32 EndFrameIndex |
| --- |
| 애니메이션 출력 시의 마지막 프레임 번호입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되어 있는지 여부를 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalPosition |
| --- |
| 이미지 출력 위치입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| 이미지 크기입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼의 Id를 지정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| float PlayRate |
| --- |
| 애니메이션 재생 속도입니다. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveSprite |
| --- |
| 원본 이미지의 비율/피봇/크기 등을 어떻게 보존하는지에 대한 타입을 정의합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| int32 StartFrameIndex |
| --- |
| 애니메이션 출력 시 시작 프레임 번호입니다. |

##### inherited from ImageComponent:

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 이미지의 기본 색상을 설정합니다 |

| boolean DropShadow |
| --- |
| 이미지의 그림자 출력 여부를 설정합니다. |

| float DropShadowAngle |
| --- |
| 그림자를 출력할 각도를 설정합니다 |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| 그림자 색상을 설정합니다. |

| float DropShadowDistance |
| --- |
| 이미지와 그림자의 거리입니다. |

| float FillAmount |
| --- |
| Type이 Filled로 설정되어 있을 때 표시되는 이미지의 비율입니다. 0부터 1 사이의 값을 사용합니다. |

| boolean FillCenter |
| --- |
| Type이 Sliced 또는 Tiled로 설정되어 있을 때 이미지 영역의 가운데를 채울지를 설정합니다. |

| boolean FillClockWise |
| --- |
| FillMethod가 Radial90, Radial180, Radial360으로 설정되어 있을 때 채우기 방향을 설정합니다. 값이 true면 시계 방향으로 채워집니다. |

| [FillMethodType](https://mod-developers.nexon.com/apiReference/Enums/FillMethodType) FillMethod |
| --- |
| Type이 Filled일 때의 채우기 방식을 설정합니다. |

| int32 FillOrigin |
| --- |
| Type이 Filled로 설정되어 있을 때 채우기 시작점을 설정합니다. FillMethod가 Horizontal 또는 Vertical일 경우 0 ~ 1 값을 사용할 수 있습니다. Radial90, Radial180, Radial360일 경우 0 ~ 3 값을 사용할 수 있습니다. |

| boolean FlipX |
| --- |
| 이미지의 X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY |
| --- |
| 이미지의 Y축을 기준으로 반전 여부를 결정합니다. |

| int32 FrameColumn ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| int32 FrameRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| int32 FrameRow ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AnimationClip Editor를 사용하세요. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ImageRUID |
| --- |
| 화면에 표시될 이미지 RUID입니다. |

| boolean Outline |
| --- |
| 이미지 외곽선 출력 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| 이미지 외곽선 색상입니다. |

| float OutlineWidth |
| --- |
| 외곽선 두께입니다. |

| boolean RaycastTarget |
| --- |
| true로 설정할 경우 화면 터치 또는 마우스 클릭 대상이 되며, 뒤에 가려진 UI는 화면 터치와 마우스 클릭 입력을 받지 못합니다. |

| [ImageType](https://mod-developers.nexon.com/apiReference/Enums/ImageType) Type |
| --- |
| 이미지를 표시하는 방식입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

##### inherited from ImageComponent:

| void SetNativeSize() |
| --- |
| 이미지를 원본 크기로 조정합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SpriteGUIAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerChangeFrameEvent) |
| --- |
| 스프라이트 애니메이션의 프레임이 바뀔 때 발생하는 이벤트입니다. |

| [SpriteGUIAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerEndEvent) |
| --- |
| 스프라이트 애니메이션 재생이 끝날 때 발생하는 이벤트입니다. |

| [SpriteGUIAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerEndFrameEvent) |
| --- |
| 지정한 값으로 SpriteAnimPlayerEndEvent를 초기화합니다. |

| [SpriteGUIAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerStartEvent) |
| --- |
| 스프라이트 애니메이션 재생이 시작될 때 발생하는 이벤트입니다. |

| [SpriteGUIAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerStartFrameEvent) |
| --- |
| 스프라이트 애니메이션이 첫번째 프레임을 재생할 때 발생하는 이벤트입니다. |

# Examples

스프라이트 애니메이션과 `Fill`을 제어하는 예제입니다. `SpriteGUIRendererComponent`가 있는 엔티티에 추가해 테스트할 수 있습니다.

```
Property:
[None]
number Fill = 0
[None]
boolean IsRequire = false

Method:
[client only]
void OnBeginPlay ()
{
	local renderer = self.Entity.SpriteGUIRendererComponent
	if not renderer then
		return
	end
	
	self.IsRequire = true
	renderer.ImageRUID = "000000"
	renderer.Type = ImageType.Filled
	renderer.FillMethod = FillMethodType.Radial360
	renderer.Color = Color.white
	renderer.AnimClipPlayType = SpriteAnimClipPlayType.ZigzagLoop
}
 
[client only]
void OnUpdate (number delta)
{
	if not self.IsRequire then
		return
	end
	
	self.Fill = math.min( self.Fill + delta, 1)
	     
	self.Entity.SpriteGUIRendererComponent.FillAmount = self.Fill
	if self.Fill >= 1.0 then
		self.Fill = 0
	end
}
```

# SeeAlso

- [FillMethodType](https://mod-developers.nexon.com/apiReference/Enums/FillMethodType)
- [ImageType](https://mod-developers.nexon.com/apiReference/Enums/ImageType)
- [SpriteAnimClipPlayType](https://mod-developers.nexon.com/apiReference/Enums/SpriteAnimClipPlayType)
- [math](https://mod-developers.nexon.com/apiReference/Lua/math)
- [Color](https://mod-developers.nexon.com/apiReference/Misc/Color)
- [기본 UI 컴포넌트](/docs?postId=744)
- [UI 제작하기](/docs?postId=64)

Update 2025-10-28 PM 02:21


# SpriteParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

스프라이트 파티클 효과를 만드는 기능을 제공합니다.

# Properties

| boolean ApplySpriteColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클이 사용할 스프라이트에 Color 프로퍼티를 적용할지 여부를 설정합니다. 프로퍼티가 false일지라도 Color의 투명도 값은 적용됩니다. |

| [SpriteParticleType](https://mod-developers.nexon.com/apiReference/Enums/SpriteParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 생성할 파티클의 타입을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클로 사용할 SpriteRUID를 설정합니다. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 반복 재생 여부를 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| 파티클을 재생합니다. |

| void Stop() |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

새 spriteRUID와 Sprite에 SpriteParticle의 Color 프로퍼티를 적용합니다. 더불어 지정한 키를 눌러 파티클 재생을 제어하는 예제입니다. Entity에 SpriteParticleComponent를 추가해야 정상 동작합니다.

```
Property:
[Sync]
SpriteParticleComponent SpriteParticle = nil

Method:
[client only]
void OnBeginPlay
{
	self.SpriteParticle = self.Entity.SpriteParticleComponent
	 
	self.SpriteParticle.SpriteRUID = "000000"
	self.SpriteParticle.ApplySpriteColor = true
}

Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: InputService
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local key = event.key
	---------------------------------------------------------
	if key == KeyboardKey.Q then
		self.SpriteParticle:Stop()
	elseif key == KeyboardKey.E then
		self.SpriteParticle:Play()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)
- [파티클 사용하기](/docs?postId=1036)
- [파티클 활용하기](/docs?postId=764)

Update 2025-12-02 PM 01:55


# SpriteRendererComponent

스프라이트 또는 애니메이션 클립을 출력합니다.

# Properties

| [Dictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/Dictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ClipName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트에 색깔을 덧씌웁니다. |

| [SpriteDrawMode](https://mod-developers.nexon.com/apiReference/Enums/SpriteDrawMode) DrawMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트가 그려지는 방식을 설정합니다. Simple, Sliced, Tiled를 사용할 수 있습니다. |

| int32 EndFrameIndex ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 재생될 애니메이션의 마지막 프레임입니다. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트의 X축을 기준으로 반전 여부를 결정합니다. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트의 Y축을 기준으로 반전 여부를 결정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 렌더러에 적용할 머티리얼 Id를 지정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 애니메이션 리소스의 경우, 재생 속도를 지정할 수 있습니다. 최소 0 이상의 값부터 지원하며 숫자가 클수록 속도가 빨라집니다. |

| RenderSettingType RenderSetting ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 스프라이트 또는 애니메이션 클립의 RUID입니다. |

| int32 StartFrameIndex ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 재생될 애니메이션의 시작 프레임입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TiledSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| DrawMode가 Tiled, Sliced일 때 스프라이트를 그릴 영역의 크기를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| 렌더러에 적용할 머티리얼을 교체합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [EmbededSpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeFrameEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerChangeFrameEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeStateEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerChangeStateEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerEndEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerEndEvent를 사용하세요. |

| [EmbededSpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerStartEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. SpriteAnimPlayerStartEvent를 사용하세요. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeFrameEvent) |
| --- |
| 스프라이트 애니메이션의 프레임이 바뀔 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeStateEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. |

| [SpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndEvent) |
| --- |
| 스프라이트 애니메이션 재생이 끝날 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndFrameEvent) |
| --- |
| 스프라이트 애니메이션이 마지막 프레임을 재생할 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartEvent) |
| --- |
| 스프라이트 애니메이션 재생이 시작될 때 발생하는 이벤트입니다. |

| [SpriteAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartFrameEvent) |
| --- |
| 스프라이트 애니메이션이 첫번째 프레임을 재생할 때 발생하는 이벤트입니다. |

# Examples

다음은 메소 아이템의 메소 값을 랜덤으로 결정하고, 값의 범위에 따라 서로 다른 `Sprite`를 적용하는 예제입니다.

```
Property:
[Sync]
number Meso = 0
 
Method:
[server only]
void OnBeginPlay ()
{
	self.Meso = _UtilLogic:RandomIntegerRange(1, 1500)
	local sprite = self.Entity.SpriteRendererComponent
	 
	if self.Meso < 50 then
		sprite.SpriteRUID = "000000"
	elseif self.Meso < 100 then
		sprite.SpriteRUID = "000000"
	elseif self.Meso < 1000 then
		sprite.SpriteRUID = "000000"
	else
		sprite.SpriteRUID = "000000"
	end
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [스프라이트 색상 조정](/docs?postId=116)
- [애니메이션 만들기](/docs?postId=595)

Update 2025-08-27 PM 04:56


# StateAnimationComponent

상태 변화에 따라 재생될 애니메이션을 지정합니다.

# Properties

| [SyncDictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 애니메이션의 이름과 AnimationClip이 매핑된 table입니다. IsLegacy의 값이 true일 때 사용됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| StateChangeEvent를 받았을 때 처리하는 함수입니다. 기본적으로 State에 매핑된 AnimationClip을 재생하는 AnimationClipEvent를 발생시킵니다. |

| void RemoveActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key) |
| --- |
| StateToAvatarBodyActionSheet에서 key에 해당하는 요소를 제거합니다. IsLegacy 값이 true면 ActionSheet에서 요소를 제거합니다. |

| void SetActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) animationClipRuid) |
| --- |
| StateToAvatarBodyActionSheet에 요소를 추가합니다. 요소로 추가되는 AvatarBodyActionElement 객체의 AvatarBodyActionStateName 프로퍼티 값은 animationClipRuid, PlayerRate 프로퍼티 값은 1이 됩니다. IsLegacy의 값이 true면 ActionSheet에 요소를 추가합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) StateStringToAnimationKey([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| State에 매핑된 Animation의 이름을 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [AnimationClipEvent](https://mod-developers.nexon.com/apiReference/Events/AnimationClipEvent) |
| --- |
| AnimationClip 변경이 필요할 때 발생하는 이벤트입니다. |

# Examples

몬스터가 무작위 피격 애니메이션을 출력하도록 하는 예제입니다. 몬스터의 기본 `StateAnimationComponent`을 삭제하고, `StateAnimationComponent`를 확장한 아래 Component를 추가합니다. 몬스터를 공격 시 다양한 주황버섯의 피격 애니메이션을 확인할 수 있습니다.

애니메이션 재생이 필요할 때 `StateStringToAnimationKey()` 함수가 호출되기 때문에 `SetRandomHitAnimation`을 미리 호출합니다. `hit` 애니메이션 키에 무작위 애니메이션 클립을 설정함으로써, `hit` 애니메이션을 재생할 때 설정된 애니메이션 클립이 재생됩니다.

```
Property:
[None]
table<string> HitAnimations

Method:
void OnBeginPlay ()
{
	table.insert(self.HitAnimations, "000000")
	table.insert(self.HitAnimations, "000000")
	table.insert(self.HitAnimations, "000000")
	table.insert(self.HitAnimations, "000000")
}

void SetRandomHitAnimation ()
{
	self:SetActionSheet("hit", self.HitAnimations[_UtilLogic:RandomIntegerRange(1, #self.HitAnimations)])
}

override string StateStringToAnimationKey (string stateName)
{
	if stateName == "HIT" then
		self:SetRandomHitAnimation()
	end

	return __base:StateStringToAnimationKey(stateName)
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [table](https://mod-developers.nexon.com/apiReference/Lua/table)

Update 2025-12-02 PM 01:55


# StateComponent

사용자 정의 StateType을 사용해 상태별 행동과 전이 규칙을 정의, 제어하는 기능을 제공합니다.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CurrentStateName ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 상태의 이름을 얻어올 수 있습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| boolean AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName, boolean reverseResult = false) |
| --- |
| stateName 상태와 nextStateName 상태를 연결합니다. 실패 시 false를 반환합니다. StateType의 OnConditionCheck의 반환 값이 true일 때 stateName 상태에서 nextStateName 상태로 상태 전이가 일어납니다. reverseResult의 값이 true면 OnConditionCheck의 반환 값이 false일 때 상태 전이가 일어납니다. |

| boolean AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName, func -> boolean conditionCheckFunction, boolean reverseResult = false) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. 다른 AddCondition(string, string, boolean) 함수를 사용하세요. |

| boolean AddState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, func updateFunction = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 함수입니다. 다른 AddState(string, Type) 함수를 사용하세요. |

| boolean AddState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, Type stateType) |
| --- |
| 사용자 정의 StateType으로 stateName이라는 이름의 상태를 추가합니다. 실패 시 false를 반환합니다. |

| boolean ChangeState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) |
| --- |
| 현재 상태를 지정한 상태로 강제 변경합니다. |

| void RemoveCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName) |
| --- |
| stateName 상태와 nextStateName 상태의 연결을 끊습니다. |

| void RemoveState([string](https://mod-developers.nexon.com/apiReference/Lua/string) name) |
| --- |
| 지정한 상태를 제거합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [DeadEvent](https://mod-developers.nexon.com/apiReference/Events/DeadEvent) |
| --- |
| 플레이어가 죽을 때 발생하는 이벤트입니다. 상태가 DEAD로 전이할 때 발생합니다. |

| [ReviveEvent](https://mod-developers.nexon.com/apiReference/Events/ReviveEvent) |
| --- |
| 플레이어가 부활할 때 발생하는 이벤트입니다. |

| [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) |
| --- |
| 상태가 변경될 때 발생하는 이벤트입니다. |

# Examples

엔티티의 현재 State를 ChatBalloon을 통해 표시해주는 컴포넌트입니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	local state = self.Entity.StateComponent
	if state == nil then
		state = self.Entity:AddComponent("StateComponent")
	end
	 
	local chatBallon = self.Entity.ChatBalloonComponent
	if chatBallon == nil then
		chatBallon = self.Entity:AddComponent("ChatBalloonComponent")
	end
	 
	 
	self.Entity.ChatBalloonComponent.AutoShowEnabled = true
	self.Entity.ChatBalloonComponent.ChatModeEnabled = false
	self.Entity.ChatBalloonComponent.ShowDuration = 1
	self.Entity.ChatBalloonComponent.HideDuration = 0
	self.Entity.ChatBalloonComponent.FontSize = 2
}

[server only]
void OnUpdate (number delta)
{
	self.Entity.ChatBalloonComponent.Message = self.Entity.StateComponent.CurrentStateName
}
```

# SeeAlso

- [ChatBalloonComponent](https://mod-developers.nexon.com/apiReference/Components/ChatBalloonComponent)
- [StateAnimationComponent](https://mod-developers.nexon.com/apiReference/Components/StateAnimationComponent)
- [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity)
- [엔티티의 상태 제어하기](/docs?postId=686)

Update 2025-08-27 PM 04:56


# StateStringToAvatarActionComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray)

더는 사용하지 않는 컴포넌트입니다. StateAnimationComponent를 사용하세요.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [BodyActionStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/BodyActionStateChangeEvent) |
| --- |
| BodyAction의 상태가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# StateStringToMonsterActionComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray)

더는 사용하지 않는 컴포넌트입니다. StateAnimationComponent를 사용하세요.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [MonsterActionStateEvent](https://mod-developers.nexon.com/apiReference/Events/MonsterActionStateEvent) |
| --- |
| 더는 사용하지 않는 이벤트입니다. StateAnimationComponent를 사용하세요. |

Update 2025-08-27 PM 04:56


# TagComponent

엔티티에 태그를 지정합니다. EntityService를 이용하면 태그를 이용해 엔티티를 조회할 수 있습니다.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Tags ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 태그 리스트를 설정합니다. 여러 개의 태그를 지정할 수 있습니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void AddTag([string](https://mod-developers.nexon.com/apiReference/Lua/string) tag) |
| --- |
| 태그를 추가합니다. |

| void RemoveTag([string](https://mod-developers.nexon.com/apiReference/Lua/string) tag) |
| --- |
| 태그를 제거합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

골인 지점 영역 안에 캐릭터가 들어오면 'Qualified' 태그를 추가합니다. 10초 뒤 해당 태그를 소유한 플레이어의 이름표를 변경해 승자를 표시하는 예제입니다.

```
Method:
[server only]
void OnBeginPlay ()
{
	local showWinner = function()
		local winners = _EntityService:GetEntitiesByTag("Qualified")
	
		for i = 1, #winners do
			local nametag = winners[i].NameTagComponent
			nametag.FontColor = Color.green
			nametag.FontSize = 4
			nametag.Name = "★ Qualified ★"
		end
	end
	 
	_TimerService:SetTimerOnce(showWinner, 10.0)
}
 
Event Handler:
[self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	 
	if TriggerBodyEntity.PlayerComponent == nil then
		return
	end
	 
	TriggerBodyEntity.TagComponent:AddTag("Qualified")
}
```

# SeeAlso

- [NameTagComponent](https://mod-developers.nexon.com/apiReference/Components/NameTagComponent)
- [PlayerComponent](https://mod-developers.nexon.com/apiReference/Components/PlayerComponent)
- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [엔티티를 탐색하는 EntityService](/docs?postId=201)

Update 2025-08-27 PM 04:56


# TextComponent

텍스트를 화면에 출력합니다. UITransformComponent와 함께 사용하는 것을 권장합니다.

# Properties

| [TextAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/TextAlignmentType) Alignment |
| --- |
| 텍스트를 정렬하는 방식입니다. |

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Text 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean BestFit |
| --- |
| 폰트 크기를 영역 크기에 맞게 조절합니다. |

| boolean Bold |
| --- |
| 굵은 텍스트 사용 여부를 설정합니다. |

| float ConstraintX |
| --- |
| 제한할 최대 너비를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| float ConstraintY |
| --- |
| 제한할 최대 높이를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| boolean DropShadow |
| --- |
| 텍스트의 그림자 출력 여부를 설정합니다. |

| float DropShadowAngle |
| --- |
| 그림자의 각도를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| 그림자의 색상입니다. |

| float DropShadowDistance |
| --- |
| 텍스트와 그림자 사이의 거리입니다. |

| [FontType](https://mod-developers.nexon.com/apiReference/Enums/FontType) Font |
| --- |
| 글꼴 종류입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor |
| --- |
| 텍스트를 렌더링하는 데 사용할 색상입니다. |

| int32 FontSize |
| --- |
| 글꼴 크기입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아와 표시합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. Text 프로퍼티의 실제 값은 변경되지 않습니다. |

| boolean IsRichText |
| --- |
| 리치 텍스트 사용 여부를 설정합니다 |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되어 있는지 여부를 나타냅니다. |

| float LineSpacing |
| --- |
| 행간 너비를 조절합니다. |

| int32 MaxSize |
| --- |
| 최대 Font 크기를 결정합니다. |

| int32 MinSize |
| --- |
| 최소 Font 크기를 결정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| 텍스트의 외곽선의 색상입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) OutlineDistance ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. OutlineWidth를 활용하세요. |

| float OutlineWidth |
| --- |
| 텍스트 외곽선의 너비를 설정합니다. |

| [OverflowType](https://mod-developers.nexon.com/apiReference/Enums/OverflowType) Overflow |
| --- |
| 텍스트가 가로 영역을 넘어가는 경우에 대한 처리 방식을 설정합니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| 텍스트 영역의 여백을 설정합니다. |

| boolean SizeFit |
| --- |
| 텍스트에 맞도록 크기를 변경합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text |
| --- |
| 보여질 내용입니다. |

| boolean UseConstraintX |
| --- |
| ConstraintX 값을 사용해 너비를 제한할 수 있도록 설정합니다. |

| boolean UseConstraintY |
| --- |
| ConstraintY 값을 사용해 높이를 제한할 수 있도록 설정합니다. |

| boolean UseNBSP |
| --- |
| 문자열이 어떻게 개행될지 설정합니다. true로 설정할 경우 문자 단위로 개행해 문자가 줄 끝에 닿으면 바로 개행합니다. false로 설정할 경우 단어 단위로 개행해 단어가 잘리지 않도록 합니다. |

| boolean UseOutLine |
| --- |
| 외곽선 효과를 사용합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| IsLocalizationKey의 값이 true인 경우 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아 반환합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. IsLocalizationKey의 값이 false인 경우 Text 프로퍼티 값을 반환합니다. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 고정된 너비의 공간에 출력되는 텍스트 영역의 높이를 가져옵니다. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 입력된 텍스트 영역의 너비를 가져옵니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

텍스트를 한 글자씩 출력하는 예제입니다.

```
Property:
[None]
number TimerID = 0
[None]
number MessageLength = 0
[None]
number MessageIdx = 0
[None]
string RawMessage = ""

Method:
[client only]
void OnBeginPlay()
{
	self.Entity.Parent.Enable = true
	 
	self.InputMessage = "안녕하세요. MSW에 오신 것을 환영합니다."
	 
	local textComponent = self.Entity.TextComponent
	if not textComponent then
		return
	end
	    
	textComponent.Bold = true
	textComponent.Alignment = TextAlignmentType.UpperLeft
	textComponent.FontColor = Color.white
	textComponent.UseOutLine = true
	textComponent.OutlineColor = Color.black
	textComponent.FontSize = 50
	 
	self:ShowDialogueOverTime(self.InputMessage, 0.1)
	-- min val is 0.1
}
 
void ShowDialogueOverTime(string rawMessage,number interval)
{
	if interval <= 0 then
		interval = 0.1
	end
	 
	self.MessageIdx = 0
	self.RawMessage = rawMessage
	self.MessageLength = utf8.len(rawMessage)
	
	local UpdateMessage = function()
		if self.MessageIdx < self.MessageLength then
			self.MessageIdx = self.MessageIdx + 1
			local currentString = _UtilLogic:SubString(self.RawMessage, 1, self.MessageIdx)
		
			if _UtilLogic:IsNilorEmptyString(currentString) == false then
				self.Entity.TextComponent.Text =  currentString
			end
		else
			self.Entity.TextComponent.Text =  ""
			self:RestartDialogue()
		end
	end
	 
	self.TimerID = _TimerService:SetTimerRepeat(UpdateMessage, interval)
}
 
void CancelDialogue()
{
	_TimerService:ClearTimer(self.TimerID)
}
 
void RestartDialogue()
{
	self.MessageIdx = 0
}
```

# SeeAlso

- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [기본 UI 컴포넌트](/docs?postId=744)
- [로컬라이징의 이해](/docs?postId=951)
- [자동 번역](/docs?postId=1072)

Update 2025-10-28 PM 02:21


# TextGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

텍스트를 화면에 출력합니다. UITransformComponent와 함께 사용하는 것을 권장합니다.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Text 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean BestFit |
| --- |
| 폰트 크기를 영역 크기에 맞게 조절합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomLeftColor |
| --- |
| 좌측 하단 색상입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomRightColor |
| --- |
| 우측 하단 색상입니다. |

| boolean ColorGradient |
| --- |
| 문자에 그라디언트 색상을 적용합니다. 그라디언트 색상은 문자 색상과 곱해집니다. |

| float ConstraintX |
| --- |
| 제한할 최대 너비를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| float ConstraintY |
| --- |
| 제한할 최대 높이를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Font |
| --- |
| 글꼴을 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor |
| --- |
| 텍스트를 렌더링하는 데 사용할 색상입니다. |

| float FontSize |
| --- |
| 글꼴 크기입니다. |

| [FontStyleType](https://mod-developers.nexon.com/apiReference/Enums/FontStyleType) FontStyle |
| --- |
| 텍스트에 적용되는 스타일입니다. |

| [GradientModes](https://mod-developers.nexon.com/apiReference/Enums/GradientModes) GradientMode |
| --- |
| 적용할 그라디언트 유형을 선택합니다. |

| [TextHorizontalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextHorizontalAlignmentOption) HorizontalAlignment |
| --- |
| 텍스트를 가로 방향으로 정렬하는 방식입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아와 표시합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. Text 프로퍼티의 실제 값은 변경되지 않습니다. |

| boolean IsRichText |
| --- |
| 리치 텍스트 사용 여부를 설정합니다. |

| float MaxSize |
| --- |
| 최대 Font 크기를 결정합니다. |

| float MinSize |
| --- |
| 최소 Font 크기를 결정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| 외곽선 색상을 설정합니다. |

| float OutlineWidth |
| --- |
| 외곽선 굵기을 설정합니다. |

| [TextOverflowMode](https://mod-developers.nexon.com/apiReference/Enums/TextOverflowMode) Overflow |
| --- |
| 텍스트가 가로 영역을 넘어가는 경우에 대한 처리 방식을 설정합니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| 텍스트 영역의 여백을 설정합니다. |

| int32 Page |
| --- |
| no description |

| boolean SizeFit |
| --- |
| 텍스트에 맞도록 크기를 변경합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [TextRendererSpacingOption](https://mod-developers.nexon.com/apiReference/Misc/TextRendererSpacingOption) SpacingOption |
| --- |
| 간격 옵션을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text |
| --- |
| 보여질 내용입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopLeftColor |
| --- |
| 좌측 상단 색상입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopRightColor |
| --- |
| 우측 상단 색상입니다. |

| boolean Underlay |
| --- |
| 그림자 생성 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) UnderlayColor |
| --- |
| 그림자 색상을 설정합니다. |

| float UnderlayOffsetX |
| --- |
| 그림자의 X축 위치를 설정합니다. |

| float UnderlayOffsetY |
| --- |
| 그림자의 Y축 위치를 설정합니다. |

| boolean UseConstraintX |
| --- |
| ConstraintX 값을 사용해 너비를 제한할 수 있도록 설정합니다. |

| boolean UseConstraintY |
| --- |
| ConstraintY 값을 사용해 높이를 제한할 수 있도록 설정합니다. |

| [TextVerticalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextVerticalAlignmentOption) VerticalAlignment |
| --- |
| 텍스트를 세로 방향으로 정렬하는 방식입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| IsLocalizationKey의 값이 true인 경우 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아 반환합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. IsLocalizationKey의 값이 false인 경우 Text 프로퍼티 값을 반환합니다. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 고정된 너비의 공간에 출력되는 텍스트 영역의 높이를 가져옵니다. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 입력된 텍스트 영역의 너비를 가져옵니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-12-02 PM 01:55


# TextGUIRendererInputComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

문자열을 입력받아 TextGUIRendererComponent에 전달합니다.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| PlaceHolder 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean AutoClear |
| --- |
| 텍스트 입력이 완료되었을 때, 입력 영역을 자동으로 초기화할지 여부를 설정합니다. |

| int32 CharacterLimit |
| --- |
| 입력 가능한 글자 수입니다. |

| [InputContentType](https://mod-developers.nexon.com/apiReference/Enums/InputContentType) ContentType |
| --- |
| 입력 가능한 타입을 지정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsFocused ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 포커싱 여부를 나타냅니다. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아와 기본 문구를 표시합니다. 이때 PlaceHolder 프로퍼티 값을 Key로 사용합니다. PlaceHolder 프로퍼티의 실제 값은 변경되지 않습니다. |

| [InputLineType](https://mod-developers.nexon.com/apiReference/Enums/InputLineType) LineType |
| --- |
| 개행 입력 방식을 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) PlaceHolder |
| --- |
| 글자가 입력되지 않은 상황에서 나타나는 기본 문구입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) PlaceHolderColor |
| --- |
| PlaceHolder 색상을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 입력한 내용입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ActivateInputField() |
| --- |
| 글자 입력을 활성화합니다. ActivateInputFied() 호출 후 몇 프레임 뒤에 IsFoucsed 값이 true로 변경됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedPlaceHolder() |
| --- |
| IsLocalizationKey의 값이 true인 경우 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아 반환합니다. 이때 PlaceHolder 프로퍼티 값을 Key로 사용합니다. IsLocalizationKey의 값이 false인 경우 PlaceHolder 프로퍼티 값을 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [TextInputEndEditEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEditorEvent) |
| --- |
| InputField의 값 변경이 완료되었을 때 발생하는 에디터 이벤트입니다. |

| [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent) |
| --- |
| InputField의 값 변경이 완료되었을 때 발생하는 이벤트입니다. |

| [TextInputSubmitEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEditorEvent) |
| --- |
| 입력을 마치고 Enter 키를 눌렀을 때 호출되는 이벤트입니다. |

| [TextInputSubmitEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEvent) |
| --- |
| 입력을 마치고 Enter 키를 눌렀을 때 호출되는 이벤트입니다. |

| [TextInputValueChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEditorEvent) |
| --- |
| InputField의 값이 변경되었을 때 발생하는 에디터 이벤트입니다. |

| [TextInputValueChangeEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEvent) |
| --- |
| InputField의 값이 변경되었을 때 발생하는 이벤트입니다. |

# Examples

아이디와 암호를 입력받는 예제입니다.

```
Property:
[None]
string Id = ""
[None]
string Password = ""
[None]
TextGUIRendererInputComponent IdInput = EntityPath
[None]
TextGUIRendererInputComponent PasswordInput = EntityPath
    
Method:
[client only]
void OnBeginPlay ()
{
	self.IdInput.Text = ""
	self.IdInput.PlaceHolder = "Id 입력"
	       
	self.PasswordInput.Text = ""
	self.PasswordInput.PlaceHolder = "패스워드 입력"
	self.PasswordInput.ContentType = InputContentType.Password
}
  
Event Handler:
[entity: EntityPath]
HandleTextInputEndEditEvent (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	log("입력된 아이디는 " .. self.InputId .. "입니다.")
}
  
[entity: EntityPath]
HandleTextInputEndEditEvent2 (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	self.Password = text
	log("입력된 패스워드는 " .. self.Password .. "입니다.")
}
```

Update 2025-12-02 PM 01:55


# TextInputComponent

문자열을 입력받아 TextComponent에 전달합니다.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| PlaceHolder 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean AutoClear |
| --- |
| 텍스트 입력이 완료되었을 때, 입력 영역을 자동으로 초기화할지 여부를 설정합니다. |

| int32 CharacterLimit |
| --- |
| 입력 가능한 글자 수입니다. |

| [InputContentType](https://mod-developers.nexon.com/apiReference/Enums/InputContentType) ContentType |
| --- |
| 입력 가능한 타입을 지정합니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsFocused ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 포커싱 여부를 나타냅니다. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아와 기본 문구를 표시합니다. 이때 PlaceHolder 프로퍼티 값을 Key로 사용합니다. PlaceHolder 프로퍼티의 실제 값은 변경되지 않습니다. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 월드에 배치되어 있는지 여부를 나타냅니다. |

| [InputLineType](https://mod-developers.nexon.com/apiReference/Enums/InputLineType) LineType |
| --- |
| 개행 입력 방식을 설정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| SortingLayer 및 OrderInLayer 값을 임의로 설정할지 여부를 결정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) PlaceHolder |
| --- |
| 글자가 입력되지 않은 상황에서 나타나는 기본 문구입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) PlaceHolderColor |
| --- |
| PlaceHolder 색상을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 입력한 내용입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void ActivateInputField() |
| --- |
| 글자 입력을 활성화합니다. ActivateInputFied() 호출 후 몇 프레임 뒤에 IsFoucsed 값이 true로 변경됩니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedPlaceHolder() |
| --- |
| IsLocalizationKey의 값이 true인 경우 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아 반환합니다. 이때 PlaceHolder 프로퍼티 값을 Key로 사용합니다. IsLocalizationKey의 값이 false인 경우 PlaceHolder 프로퍼티 값을 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [TextInputEndEditEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEditorEvent) |
| --- |
| InputField의 값 변경이 완료되었을 때 발생하는 에디터 이벤트입니다. |

| [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent) |
| --- |
| InputField의 값 변경이 완료되었을 때 발생하는 이벤트입니다. |

| [TextInputSubmitEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEditorEvent) |
| --- |
| 입력을 마치고 Enter 키를 눌렀을 때 호출되는 이벤트입니다. |

| [TextInputSubmitEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEvent) |
| --- |
| 입력을 마치고 Enter 키를 눌렀을 때 호출되는 이벤트입니다. |

| [TextInputValueChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEditorEvent) |
| --- |
| InputField의 값이 변경되었을 때 발생하는 에디터 이벤트입니다. |

| [TextInputValueChangeEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEvent) |
| --- |
| InputField의 값이 변경되었을 때 발생하는 이벤트입니다. |

# Examples

아이디와 패스워드를 입력받는 예제입니다.

```
Property:
[None]
string Id = ""
[None]
string Password = ""
[None]
TextInputComponent IdInput = EntityPath
[None]
TextInputComponent PasswordInput = EntityPath
 
 
Method:
[client only]
void OnBeginPlay ()
{
	self.IdInput.Text = ""
	self.IdInput.PlaceHolder = "Id 입력"
	  
	self.PasswordInput.Text = ""
	self.PasswordInput.PlaceHolder = "패스워드 입력"
	self.PasswordInput.ContentType = InputContentType.Password
}
 
Event Handler:
[entity: EntityPath]
HandleTextInputEndEditEvent (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	log("입력된 아이디는 " .. self.InputId .. "입니다.")
}
 
[entity: EntityPath]
HandleTextInputEndEditEvent2 (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	self.Password = text
	log("입력된 패스워드는 " .. self.Password .. "입니다.")
}
```

# SeeAlso

- [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent)
- [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity)
- [기본 UI 컴포넌트](/docs?postId=744)

Update 2025-10-28 PM 02:21


# TextRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

월드 공간에 텍스트를 출력합니다. TransformComponent와 함께 사용합니다.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Text 프로퍼티 값의 자동 번역 여부를 설정합니다. |

| boolean BestFit ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 폰트 크기를 영역 크기에 맞게 조절합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomLeftColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 하단 촤측 색상입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomRightColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 하단 우측 색상입니다. |

| boolean ColorGradient ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 문자에 그라이언트 색상을 적용합니다. 문자 색상은 적용한 그라이던트 색상과 곱해져 보여집니다. |

| float ConstraintX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 제한할 최대 너비를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| float ConstraintY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 제한할 최대 높이를 설정합니다. Sizefit이 true인 경우 동작합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Font ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 글꼴을 설정합니다 |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트를 렌더링하는 데 사용할 색상입니다. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 글꼴 크기입니다. |

| [FontStyleType](https://mod-developers.nexon.com/apiReference/Enums/FontStyleType) FontStyle ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트에 적용되는 스타일입니다. |

| GradientType GradientMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 적용할 그라디언트 유형을 선택합니다. |

| [TextHorizontalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextHorizontalAlignmentOption) HorizontalAlignment ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 가로 방향 텍스트를 정렬하는 방식입니다. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아와 표시합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. Text 프로퍼티의 실제 값은 변경되지 않습니다. |

| boolean IsRichText ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 리치 텍스트 사용 여부를 설정합니다. |

| float MaxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 최대 폰트 크기를 결정합니다. |

| float MinSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 최소 폰트 크기를 결정합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 외곽선 색상을 설정합니다. |

| float OutlineWidth ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 외곽선의 굵기를 설정합니다. |

| [TextOverflowMode](https://mod-developers.nexon.com/apiReference/Enums/TextOverflowMode) Overflow ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트 박스의 가로 영역을 넘어가는 텍스트 처리 방식을 설정합니다. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트 영역의 여백을 설정합니다. |

| int32 Page ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트를 여러 장으로 잘라 표시합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RectOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트 박스의 기준점입니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RectSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트 박스의 사이즈입니다. |

| boolean SizeFit ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 텍스트에 맞도록 텍스트 박스 크기를 변경합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| [TextRendererSpacingOption](https://mod-developers.nexon.com/apiReference/Misc/TextRendererSpacingOption) SpacingOption ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 간격 옵션을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 보여질 내용입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopLeftColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 상단 좌측 색상입니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopRightColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 상단 우측 색상입니다. |

| boolean Underlay ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 글자의 그림자 생성 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) UnderlayColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자 색상을 설정합니다. |

| float UnderlayOffsetX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자의 X축 위치를 설정합니다. |

| float UnderlayOffsetY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 그림자의 Y축 위치를 설정합니다. |

| boolean UseConstraintX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| ConstraintX 값을 사용해 너비를 제한할 수 있도록 설정합니다. |

| boolean UseConstraintY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| ConstraintY 값을 사용해 높이를 제한할 수 있도록 설정합니다. |

| [TextVerticalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextVerticalAlignmentOption) VerticalAlignment ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 세로 방향 텍스트를 정렬하는 방식입니다. |

| boolean Wrapping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 줄바꿈 유무 여부를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| IsLocalizationKey의 값이 true인 경우 LocaleDataSet에서 현재 언어 설정에 맞는 텍스트를 찾아 반환합니다. 이때 Text 프로퍼티 값을 Key로 사용합니다. IsLocalizationKey의 값이 false인 경우 Text 프로퍼티 값을 반환합니다. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 고정된 너비의 공간에 출력되는 텍스트 영역의 높이를 가져옵니다. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| 입력된 텍스트 영역의 너비를 가져옵니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| 컴포넌트의 SortingLayer가 변경되었을 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# TileMapComponent

메이플스토리 지형 기능을 제공합니다. 맵 레이어당 하나만 존재할 수 있습니다.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 타일맵의 색상을 지정합니다. |

| boolean CreateFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 발판의 생성 여부를 결정합니다. false일 시 발판은 생성되지 않습니다. |

| float FootholdDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 타일맵 위에 있는 플레이어의 마찰력을 변화시킵니다. 값이 클수록 더욱 더 빠르게 감속합니다. |

| float FootholdForce ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 타일맵 위에 있는 플레이어에 가하는 힘입니다. 양수일 경우 오른쪽, 음수일 경우 왼쪽으로 힘을 받습니다. |

| float FootholdWalkSpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 타일맵 위에 있는 플레이어의 속력을 변화시킵니다. 값이 클수록 더욱 더 빨라집니다. |

| boolean IgnoreMapLayerCheck |
| --- |
| SortingLayer에 Map Layer 이름을 지정했을 때 자동 치환을 수행하지 않습니다. |

| boolean IncludeFinishFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| 발판을 생성할 때 발판의 굽은 끝 부분 포함 여부를 결정합니다. |

| boolean IsBlockVerticalLine |
| --- |
| 엔티티가 해당 타일맵의 세로 발판에 막힐지 결정합니다. RigidbodyComponent를 가진 엔티티만 해당합니다. |

| boolean IsOddGridPosition |
| --- |
| 타일맵을 그리드의 기준점과 어긋나게 배치합니다. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 같은 Layer 내의 우선순위를 결정합니다. 수가 클수록 앞에 보입니다. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| true일 경우 Physics 기능을 사용하는 Dynamic 강체(PhysicRigidbody)와 충돌할 수 있습니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 2개 이상의 엔티티가 겹쳤을 때 Sorting Layer에 따라 보이는 우선순위가 결정됩니다. |

| TileMapVersion TileMapVersion ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 타일맵 생성 규칙 버전입니다. 과거 버전과의 호환성을 위해 존재하는 프로퍼티입니다. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) TileSetRUID |
| --- |
| 타일맵에서 사용할 타일셋의 RUID를 지정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| Component의 OrderInLayer가 변경되었을 때 발생하는 이벤트입니다. |

# Examples

캐릭터가 엔티티에 출동하면 타일 맵의 색상이 변경됩니다. 충돌 범위를 벗어나면 본래 색상으로 돌아오게 됩니다.

```
Method:
void SetColor (Color color)
{
	local tileMap = _EntityService:GetEntityByPath("/maps/map01/TileMap")
	tileMap.TileMapComponent.Color = color
}

Event Handler:
[self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	if TriggerBodyEntity.Name == "Name" then
		self:SetColor(Color.red)
	end
}

[self]
HandleTriggerLeaveEvent (TriggerLeaveEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	self:SetColor(Color.white)
}
```

# SeeAlso

- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent)
- [Color](https://mod-developers.nexon.com/apiReference/Misc/Color)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)

Update 2025-08-27 PM 04:56


# TouchReceiveComponent

엔티티를 터치할 수 있게 되며 터치 시 동작을 제어할 수 있습니다.

# Properties

| boolean AutoFitOnce ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| true로 설정하면 AvatarRendererComponent 또는 SpriteRendererComponent의 Scale 값에 맞추어 TouchArea 및 Offset 값을 자동으로 변경됩니다. 최초 한 번만 변경합니다. |

| boolean AutoFitToSize |
| --- |
| true로 설정하면 AvatarRendererComponent 또는 SpriteRendererComponent의 Scale 값이 변할 때마다 TouchArea와 Offset 값이 자동으로 변경됩니다. |

| boolean DynamicTouchArea ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AutoFitToSize 프로퍼티를 사용하세요. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Offset |
| --- |
| 터치 영역의 중심점을 설정합니다. |

| boolean RelayEventToBehind |
| --- |
| 렌더링 순서상 뒤에 위치한 TouchReceiveComponent의 터치 이벤트 발생 여부를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TouchArea |
| --- |
| 터치 영역의 크기입니다. AutoFitToSize가 true일 경우 중간에 TouchArea의 값이 변경될 수 있습니다. |

| float TouchAreaUpdateTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. AutoFitToSize 프로퍼티를 사용하세요. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [TouchEvent](https://mod-developers.nexon.com/apiReference/Events/TouchEvent) |
| --- |
| 엔티티를 터치했을 때 발생하는 이벤트입니다. 엔티티에 TouchReceiveComponent가 있어야 이벤트가 발생합니다. |

| [TouchHoldEvent](https://mod-developers.nexon.com/apiReference/Events/TouchHoldEvent) |
| --- |
| 엔티티를 터치하는 동안 발생하는 이벤트입니다. 엔티티에 TouchReceiveComponent가 있어야 이벤트가 발생합니다. 짧게 터치하면 발생하지 않습니다. |

| [TouchReleaseEvent](https://mod-developers.nexon.com/apiReference/Events/TouchReleaseEvent) |
| --- |
| 엔티티를 터치를 뗐을 때 발생하는 이벤트입니다. 엔티티에 TouchReceiveComponent가 있어야 이벤트가 발생합니다. 짧게 터치하면 발생하지 않습니다. |

# Examples

#### TouchEvent

다음은 NPC를 터치하면 관련 UI를 표시하는 예제입니다. NPC와의 대화, 거래 등에 응용할 수 있습니다.

```
Event Handler:
[self]
HandleTouchEvent (TouchEvent event)
{
	-- Parameters
	--------------------------------------------------------
	 
	local uiPath = "/ui/DefaultGroup/ExampleUI"
	local targetUI = _EntityService:GetEntityByPath(uiPath)
	 
	targetUI:SetVisible(true)
}
```

#### TouchHoldEvent

플레이어가 엔티티를 터치하고 이동시키는 예제입니다. `TouchReceiveComponent`가 엔티티에 추가되어야 합니다.

```
Property:
[None]
table<integer> Index
[None]
any Handler = nil

Method:
[client]
void Move (any screenholdEvent)
{
	local touchId = screenholdEvent.TouchId
	local touchPoint = screenholdEvent.TouchPoint
	 
	local string
	for i =1, #self.Index do
		if self.Index[i] == touchId then
			self.Entity.TransformComponent.WorldPosition = _UILogic:ScreenToWorldPosition(touchPoint):ToVector3()
			break
		end
	end
}

[client]
void AddHandler (integer touchId)
{
	if self.Handler ~= nil then
		return
	end
	 
	table.insert(self.index, touchId)
	self.Handler = _InputService:ConnectEvent(ScreenTouchHoldEvent, self.Move)
}

Event Handler:
[self]
HandleTouchEvent (TouchEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TouchReceiveComponent
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local TouchId = event.TouchId
	local TouchPoint = event.TouchPoint
	---------------------------------------------------------
	 
	self:AddHandler(TouchId)
}

[self]
HandleTouchHoldEvent (TouchHoldEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TouchReceiveComponent
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local TouchId = event.TouchId
	local TouchPoint = event.TouchPoint
	---------------------------------------------------------
	self:AddHandler(TouchId)
}

[self]
HandleTouchReleaseEvent (TouchReleaseEvent event)
{
	--------------- Native Event Sender Info ----------------
	-- Sender: TouchReceiveComponent
	-- Space: Client
	---------------------------------------------------------
	 
	-- Parameters
	local TouchId = event.TouchId
	local TouchPoint = event.TouchPoint
	---------------------------------------------------------
	for i =1, #self.Index do
		if self.Index[i] == TouchId then
			table.remove(self.index, i)
			if #self.Index <= 0 then
				_InputService:DisconnectEvent(ScreenTouchHoldEvent, self.handler)
				self.Handler = nil
			end
			break
		end
	end
}
```

# SeeAlso

- [TransformComponent](https://mod-developers.nexon.com/apiReference/Components/TransformComponent)
- [TouchEvent](https://mod-developers.nexon.com/apiReference/Events/TouchEvent)
- [UILogic](https://mod-developers.nexon.com/apiReference/Logics/UILogic)
- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)

Update 2025-08-27 PM 04:56


# TransformComponent

엔티티의 위치, 회전, 크기를 나타냅니다.

# Properties

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Position ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티의 부모를 기준으로 좌표를 나타냅니다. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) QuaternionRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 엔티티의 회전 값을 Quaternion으로 나타냅니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Rotation |
| --- |
| 엔티티의 회전 값을 오일러 각으로 나타냅니다. 2차원에서는 Z축을 이용하여 회전할 수 있습니다. QuaternitonRotation에 의해 동기화됩니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Scale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티의 크기 비율을 나타냅니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) WorldPosition |
| --- |
| 엔티티의 월드 기준 좌표를 나타냅니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) WorldRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 엔티티의 월드 기준 회전 값을 오일러 각으로 나타냅니다. |

| float WorldZRotation |
| --- |
| 엔티티의 월드 기준 오일러 각 회전 값 중 Z축의 값을 나타냅니다. |

| float ZRotation |
| --- |
| 엔티티의 오일러 각 회전 값 중 Z축의 값을 나타냅니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| [FastVector3](https://mod-developers.nexon.com/apiReference/Misc/FastVector3) PositionAsFastVector3() |
| --- |
| Position값을 FastVector3 타입으로 반환합니다. |

| void Rotate(float angle) |
| --- |
| 이 엔티티를 angle만큼 반시계 방향으로 회전시킵니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToLocalDirection([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldDirection) |
| --- |
| 입력 받은 방향을 월드 좌표에서 로컬 좌표로 변환합니다. Scale과 Position에 영향을 받지 않습니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToLocalPoint([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPoint) |
| --- |
| 입력 받은 월드 좌표를 로컬 좌표로 변환합니다. Scale에 영향을 받습니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldDirection([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) localDirection) |
| --- |
| 입력 받은 방향을 로컬 좌표에서 월드 좌표로 변환합니다. Scale과 Position에 영향을 받지 않습니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPoint([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) localPoint) |
| --- |
| 입력 받은 로컬 좌표를 월드 좌표로 변환합니다. Scale에 영향을 받습니다. |

| void Translate(float deltaX, float deltaY) |
| --- |
| 이 엔티티의 좌표를 deltaX, deltaY만큼 이동시킵니다. |

| [FastVector3](https://mod-developers.nexon.com/apiReference/Misc/FastVector3) WorldPositionAsFastVector3() |
| --- |
| WorldPosition 값을 FastVector3 타입으로 반환합니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

다음은 `ZRotation` 프로퍼티를 이용해 Entity를 일정한 속력으로 회전시키는 예제입니다.

```
Property:
[None]
number AngularSpeed = 360
 
Method:
[server only]
void OnUpdate (number delta)
{
	local transform = self.Entity.TransformComponent
	local zRotation = transform.ZRotation

	transform.ZRotation = zRotation + (self.AngularSpeed * delta)
}
```

동일한 동작을 하는 코드를 `Rotate` 함수를 이용해 작성하면 다음과 같습니다.

```
Property:
[None]
number AngularSpeed = 360
 
Method:
[server only]
void OnUpdate (number delta)
{
	local transform = self.Entity.TransformComponent
	 
	transform:Rotate(self.AngularSpeed * delta)	
}
```

다음은 엔티티를 자유낙하 시키는 예제입니다. `Translate` 함수와 `delta`를 이용해 Entity를 현재 속도만큼 이동시킵니다.

```
Property:
[None]
Vector2 Gravity = Vector2(0,-9.8)
[Sync]
Vector2 CurrentVelocity = Vector2(0,0)

Method:
[server only]
void OnUpdate (number delta)
{
	local transform = self.Entity.TransformComponent
 
	self.CurrentVelocity = self.CurrentVelocity + (self.Gravity * delta)
 
	local deltaX = self.CurrentVelocity.x * delta
	local deltaY = self.CurrentVelocity.y * delta
 
	transform:Translate(deltaX, deltaY)
}
```

# SeeAlso

- [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2)
- [엔티티의 위치, 크기, 회전 조정](/docs?postId=82)

Update 2025-08-27 PM 04:56


# TriggerComponent

엔티티에 충돌 영역을 설정하고 충돌을 감지할 수 있는 기능을 제공합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| IsLegacy가 true인 이전 시스템에서 사용할 수 있습니다. 엔티티를 기준으로 충돌체 직사각형의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 직사각형 충돌체의 너비와 높이를 지정합니다. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 원형 충돌체의 반지름입니다. ColliderType이 Circle일 때 유효합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 더는 사용하지 않는 프로퍼티입니다. CollisionGroup을 사용하세요. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 엔티티를 기준으로 충돌체의 중심점 위치를 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌체의 타입을 설정합니다. IsLegacy가 false인 신규 시스템에서 사용할 수 있습니다. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 충돌 그룹을 설정합니다. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트가 이전의 시스템으로 동작할지를 설정합니다. 신규 시스템은 충돌체가 TransformComponent의 회전과 크기에 영향을 받습니다. 또한 ColliderType을 설정해 원 모양 충돌체를 사용할 수 있습니다. |

| boolean IsPassive ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true인 경우 스스로 충돌 검사를 하지 않습니다. IsPassive가 true인 TriggerComponent끼리 충돌하면 이벤트가 발생하지 않습니다. 이벤트가 발생하기 위해선 적어도 TriggerComponent 중 하나는 IsPassive가 false여야 합니다. 불필요한 충돌 검사를 줄여 월드 성능을 개선할 수 있습니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 다각형 충돌체를 이루는 점들의 위치입니다. ColliderType이 Polygon일 때 유효합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void OnEnterTriggerBody([TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent) enterEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 트리거 영역에 진입했을 때 호출되는 함수입니다. |

| void OnLeaveTriggerBody([TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent) leaveEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 트리거 영역을 벗어날 때 호출되는 함수입니다. |

| void OnStayTriggerBody([TriggerStayEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerStayEvent) stayEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| 트리거 영역에 진입해서 남아있을 때 매 프레임마다 호출되는 함수입니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent) |
| --- |
| TriggerComponent의 영역이 겹치는 순간 발생하는 이벤트입니다. |

| [TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent) |
| --- |
| 겹쳐 있던 TriggerComponent의 영역이 떨어지는 순간 발생하는 이벤트입니다. |

| [TriggerStayEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerStayEvent) |
| --- |
| TriggerComponent의 영역이 겹쳐 있는 동안 매 프레임마다 발생하는 이벤트입니다. 월드 성능 하락의 요인이 될 수 있으므로 사용 시 주의가 필요합니다. |

# Examples

다음은 체력 회복 공간을 만드는 예제입니다. 플레이어가 `Heal`이라는 이름의 Entity의 트리거 영역 안에 있는 동안 Hp를 증가시킵니다. 플레이어가 트리거 영역 안에 들어와 있는지 알기 위해 `TriggerEnterEvent`, `TriggerLeaveEvent`를 사용할 수 있습니다.

```
Property:
[Sync]
boolean IsGettingHealed = false
[Sync]
number Hp = 1000
 
Method:
[server only]
void OnUpdate (number delta)
{
	if self.IsGettingHealed then
		local player = self.Entity.PlayerComponent
		     
		self.Hp = self.Hp + (10.0 * delta)
		player.Hp = math.floor(self.Hp)
	end
}
 
Event Handler:
[server only] [self]
HandleTriggerEnterEvent (TriggerEnterEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	 
	if TriggerBodyEntity.Name == "Heal" then
		self.IsGettingHealed = true
	end
}
 
[server only] [self]
HandleTriggerLeaveEvent (TriggerLeaveEvent event)
{
	-- Parameters
	local TriggerBodyEntity = event.TriggerBodyEntity
	--------------------------------------------------------
	 
	if TriggerBodyEntity.Name == "Heal" then
		self.IsGettingHealed = false
	end
}
```

# SeeAlso

- [PlayerComponent](https://mod-developers.nexon.com/apiReference/Components/PlayerComponent)
- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent)
- [math](https://mod-developers.nexon.com/apiReference/Lua/math)

Update 2025-08-27 PM 04:56


# TweenBaseComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Abstract&amp;color=darkkhaki)

Tween 종류의 부모 컴포넌트입니다.

# Properties

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우, 트윈이 목적지에 도달하여 끝나거나 직접 Stop()을 호출했을 때 이 컴포넌트를 엔티티에서 제거합니다. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우, 게임이 시작되면 트윈이 자동으로 재생합니다. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 재생 상태입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 위치입니다. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 회전값입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 크기입니다. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 재생 주체를 설정합니다. 기본값은 Default입니다. 서버와 클라이언트에서 함수로 트윈을 제어할 수 있습니다. 서버 상태는 클라이언트로 항상 동기화됩니다. 실행 공간 제어를 client only로 설정하면 클라이언트에서만 트윈을 제어하도록 제한하고, 서버로 동기화하지 않습니다.<br>트윈 제어 함수: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈이 시작된 후 경과한 시간입니다. 초 단위입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void Destroy() |
| --- |
| 컴포넌트를 제거합니다. CurrentState를 Destroying으로 변경합니다. |

| void Pause() |
| --- |
| 트윈을 일시정지합니다. CurrentState를 Pausing으로 변경합니다. |

| void Play() |
| --- |
| 트윈을 재생합니다. CurrentState를 Playing으로 변경합니다. |

| void RestartFromCurrentPosition() |
| --- |
| 현재 위치에서 재시작합니다. |

| void Stop(boolean reset) |
| --- |
| 트윈을 중지합니다. CurrentState를 Idle로 변경합니다. reset에 true를 지정할 경우 트윈의 위치와 진행도를 초기 상태로 되돌립니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# TweenCircularComponent

Entity가 원 경로를 따라 회전합니다.

# Properties

| float Degree ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 총 회전 각도를 설정합니다. 값이 360인 경우 한 바퀴 회전하고, 0인 경우 무한히 회전합니다. |

| boolean IsClockwise ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 어느 방향으로 회전할지 정의합니다. true인 경우 시계 방향으로 회전합니다. |

| boolean LookAtCenter ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true인 경우 중심을 바라보며 회전합니다. |

| float Radius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 회전 반경을 설정합니다. |

| float Speed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 초당 몇 도만큼 회전할지 회전 속도를 설정합니다. 값이 360인 경우 1초에 한 바퀴 회전합니다. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우, 트윈이 목적지에 도달하여 끝나거나 직접 Stop()을 호출했을 때 이 컴포넌트를 엔티티에서 제거합니다. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우, 게임이 시작되면 트윈이 자동으로 재생합니다. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 재생 상태입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 위치입니다. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 회전값입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 크기입니다. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 재생 주체를 설정합니다. 기본값은 Default입니다. 서버와 클라이언트에서 함수로 트윈을 제어할 수 있습니다. 서버 상태는 클라이언트로 항상 동기화됩니다. 실행 공간 제어를 client only로 설정하면 클라이언트에서만 트윈을 제어하도록 제한하고, 서버로 동기화하지 않습니다.<br>트윈 제어 함수: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈이 시작된 후 경과한 시간입니다. 초 단위입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| 컴포넌트를 제거합니다. CurrentState를 Destroying으로 변경합니다. |

| void Pause() |
| --- |
| 트윈을 일시정지합니다. CurrentState를 Pausing으로 변경합니다. |

| void Play() |
| --- |
| 트윈을 재생합니다. CurrentState를 Playing으로 변경합니다. |

| void RestartFromCurrentPosition() |
| --- |
| 현재 위치에서 재시작합니다. |

| void Stop(boolean reset) |
| --- |
| 트윈을 중지합니다. CurrentState를 Idle로 변경합니다. reset에 true를 지정할 경우 트윈의 위치와 진행도를 초기 상태로 되돌립니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

`TweenCircularComponent`가 있는 엔티티를 무한히 회전하게 만드는 예제입니다.

```
local tween =self.Entity.TweenCircularComponent
tween.Degree = 0

-- 회전 반경을 2로 설정합니다.
tween.Radius = 2

-- 1초에 180도만큼의 속력으로 회전합니다. 한 바퀴 회전하는 데 2초가 소요됩니다.
tween.Speed = 180
```

# SeeAlso

- [엔티티 구간 이동시키기](/docs?postId=122)

Update 2025-08-27 PM 04:56


# TweenFloatingComponent

Entity가 위 아래로 왕복 이동합니다.

# Properties

| float Amplitude ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 최저점에서 최고점까지의 거리를 설정합니다. 값이 1인 경우 시작 위치에서 최고/최저점까지의 거리는 0.5가 됩니다. |

| float OneCycleTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 최저점에서 최고점까지 이동하는데 걸리는 시간을 설정합니다. 초 단위입니다.<br>예를 들어, 값이 1인 경우 한 이동 주기는 2초가 소요됩니다. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 움직임 효과를 설정합니다. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우, 트윈이 목적지에 도달하여 끝나거나 직접 Stop()을 호출했을 때 이 컴포넌트를 엔티티에서 제거합니다. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우, 게임이 시작되면 트윈이 자동으로 재생합니다. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 재생 상태입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 위치입니다. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 회전값입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 크기입니다. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 재생 주체를 설정합니다. 기본값은 Default입니다. 서버와 클라이언트에서 함수로 트윈을 제어할 수 있습니다. 서버 상태는 클라이언트로 항상 동기화됩니다. 실행 공간 제어를 client only로 설정하면 클라이언트에서만 트윈을 제어하도록 제한하고, 서버로 동기화하지 않습니다.<br>트윈 제어 함수: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈이 시작된 후 경과한 시간입니다. 초 단위입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| 컴포넌트를 제거합니다. CurrentState를 Destroying으로 변경합니다. |

| void Pause() |
| --- |
| 트윈을 일시정지합니다. CurrentState를 Pausing으로 변경합니다. |

| void Play() |
| --- |
| 트윈을 재생합니다. CurrentState를 Playing으로 변경합니다. |

| void RestartFromCurrentPosition() |
| --- |
| 현재 위치에서 재시작합니다. |

| void Stop(boolean reset) |
| --- |
| 트윈을 중지합니다. CurrentState를 Idle로 변경합니다. reset에 true를 지정할 경우 트윈의 위치와 진행도를 초기 상태로 되돌립니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

Entity가 부유하는 듯한 효과를 주는 예제입니다.

```
local tween = self.Entity.TweenFloatingComponent
 
-- 이동 거리 0.1
tween.Amplitude = 0.1
-- 주기 1초
tween.OneCycleTime = 1
-- 서서히 느려지다가 다시 빨라지는 효과
tween.TweenType = EaseType.QuadEaseInOut
```

# SeeAlso

- [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType)

Update 2025-08-27 PM 04:56


# TweenLineComponent

엔티티가 지정한 경로로 이동합니다. 다양한 움직임 효과를 지원합니다.

# Properties

| [CoordinateType](https://mod-developers.nexon.com/apiReference/Enums/CoordinateType) DestinationCoordinateType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Positions가 절대 좌표인지, 상대 좌표인지 설정합니다. 재생 중 변경되면 트윈을 재시작합니다. |

| float Duration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 출발지에서 목적지까지 편도 이동 시간을 설정합니다. 초 단위로 설정합니다.<br>예를 들어, Duration이 1이고 OneRoundTrip으로 이동한다면 각 편도 이동 시간은 정방향 1초, 역방향 1초이며 총 이동 시간은 2초입니다. |

| [InterpolationType](https://mod-developers.nexon.com/apiReference/Enums/InterpolationType) Interpolation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Positions에서 설정한 이동 경로의 보간 방식을 설정합니다. 재생 중 변경되면 트윈을 재시작합니다. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Positions ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 이동 경로를 설정합니다. 재생 중 변경되면 트윈을 재시작합니다. |

| float ReturnDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 목적지에서 출발지로 되돌아올 때 걸리는 시간을 설정합니다. 초 단위로 설정합니다. UseReturnTweenType이 true일 때 유효합니다. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) ReturnTweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 목적지에서 출발지로 되돌아 올 때의 움직임 효과를 설정합니다. UseReturnTweenType이 true일 때 유효합니다. |

| [TweenLinearStopType](https://mod-developers.nexon.com/apiReference/Enums/TweenLinearStopType) StopType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 편도 또는 왕복 이동 여부를 설정합니다. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 움직임 효과를 설정합니다. |

| boolean UseReturnTweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| ReturnDuration과 ReturnTweenType 적용 여부를 설정합니다. 값이 false인 경우 목적지에서 출발지로 되돌아 올 때도 Duration과 TweenType이 적용됩니다. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| true일 경우, 트윈이 목적지에 도달하여 끝나거나 직접 Stop()을 호출했을 때 이 컴포넌트를 엔티티에서 제거합니다. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 값이 true일 경우, 게임이 시작되면 트윈이 자동으로 재생합니다. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 현재 재생 상태입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 위치입니다. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 회전값입니다. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈의 시작 크기입니다. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 재생 주체를 설정합니다. 기본값은 Default입니다. 서버와 클라이언트에서 함수로 트윈을 제어할 수 있습니다. 서버 상태는 클라이언트로 항상 동기화됩니다. 실행 공간 제어를 client only로 설정하면 클라이언트에서만 트윈을 제어하도록 제한하고, 서버로 동기화하지 않습니다.<br>트윈 제어 함수: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 트윈이 시작된 후 경과한 시간입니다. 초 단위입니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| 컴포넌트를 제거합니다. CurrentState를 Destroying으로 변경합니다. |

| void Pause() |
| --- |
| 트윈을 일시정지합니다. CurrentState를 Pausing으로 변경합니다. |

| void Play() |
| --- |
| 트윈을 재생합니다. CurrentState를 Playing으로 변경합니다. |

| void RestartFromCurrentPosition() |
| --- |
| 현재 위치에서 재시작합니다. |

| void Stop(boolean reset) |
| --- |
| 트윈을 중지합니다. CurrentState를 Idle로 변경합니다. reset에 true를 지정할 경우 트윈의 위치와 진행도를 초기 상태로 되돌립니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Examples

Tween 움직임을 키 입력에 따라 정지, 실행하는 예시 코드입니다.

```
Method:
[server]
void StopOrPlay ()
{
	local tween = self.Entity.TweenLineComponent
	 
	if tween.CurrentState == TweenState.Playing then
		tween:Pause()
	else
		tween:Play()
	end
}

Event Handler:
[service: InputService]
HandleKeyDownEvent (KeyDownEvent event)
{
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.Q then
		self:StopOrPlay()
	end
}
```

# SeeAlso

- [KeyDownEvent](https://mod-developers.nexon.com/apiReference/Events/KeyDownEvent)

Update 2025-08-27 PM 04:56


# UIAreaParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

UI에 생성 영역을 조절할 수 있는 파티클 효과를 만드는 기능을 제공합니다.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaOffset |
| --- |
| 엔티티를 기준으로 생성 범위의 중심점 위치를 설정합니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaSize |
| --- |
| 파티클 생성 범위의 너비와 높이를 지정합니다. |

| [UIAreaParticleType](https://mod-developers.nexon.com/apiReference/Enums/UIAreaParticleType) ParticleType |
| --- |
| 생성할 파티클 타입을 설정합니다. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| 파티클의 크기입니다. |

| boolean Loop |
| --- |
| 파티클의 반복 재생 여부를 설정합니다. |

| float ParticleCount |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클을 재생합니다. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# UIBaseParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)![custom](https://img.shields.io/static/v1?label=&amp;message=Abstract&amp;color=darkkhaki)

UI에 사용하는 파티클 효과를 만드는 UIParticleComponent의 부모 컴포넌트입니다.

# Properties

| boolean AutoRandomSeed |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| 파티클의 크기입니다. |

| boolean Loop |
| --- |
| 파티클의 반복 재생 여부를 설정합니다. |

| float ParticleCount |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클을 재생합니다. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# UIBasicParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

UI에 사용하는 기본 파티클의 설정 및 제어 기능을 제공합니다.

# Properties

| [UIBasicParticleType](https://mod-developers.nexon.com/apiReference/Enums/UIBasicParticleType) ParticleType |
| --- |
| 생성할 파티클의 타입을 설정합니다. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| 파티클의 크기입니다. |

| boolean Loop |
| --- |
| 파티클의 반복 재생 여부를 설정합니다. |

| float ParticleCount |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클을 재생합니다. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56


# UIGroupComponent

UIGroup을 나타냅니다. UI Entity를 그룹화하고 그룹의 속성을 설정할 수 있습니다. UIGroup은 UIGroup 편집 창에서 생성하고 삭제할 수 있습니다.

# Properties

| boolean DefaultShow |
| --- |
| 게임을 시작했을 때 그룹의 활성화 여부를 설정합니다. false로 설정할 경우 비활성화된 상태로 시작합니다. |

| int32 GroupOrder |
| --- |
| UIGroup 간 레이어 순서입니다. UIGroup 편집 창에서는 GroupOrder 값이 클수록 위에 표시됩니다. |

| [UIGroupType](https://mod-developers.nexon.com/apiReference/Enums/UIGroupType) GroupType |
| --- |
| UIGroup의 타입을 설정합니다. <br><br>* DefaultType: 기본 UIGroup입니다. 월드 최초 생성 시 자동으로 생성되고, 삭제할 수 없습니다.<br>* UIType: UI 에디터에서 직접 생성한 UIGroup입니다.<br>* EditorType: UI 에디터에서 직접 생성한 UIGroup입니다. Scene 편집 중 사용할 수 있는 Editor UIGroup이 됩니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

Update 2025-08-27 PM 04:56


# UISpriteParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

기본 파티클의 설정 및 제어 기능을 제공합니다.

# Properties

| boolean ApplySpriteColor |
| --- |
| 파티클이 사용할 스프라이트에 Color 프로퍼티를 적용할지 여부를 설정합니다. 프로퍼티가 false일지라도 Color의 투명도 값은 적용됩니다. |

| [UISpriteParticleType](https://mod-developers.nexon.com/apiReference/Enums/UISpriteParticleType) ParticleType |
| --- |
| 생성할 파티클의 타입을 설정합니다. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID |
| --- |
| 파티클로 사용할 SpriteRUID를 설정합니다. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| 파티클 입자 방출이 시작될 때마다 랜덤 시드를 새로 생성할지 여부를 설정합니다. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| 렌더링될 파티클의 색상을 보정합니다. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 새 파티클을 방출하고 있는지를 나타냅니다. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| 파티클의 크기입니다. |

| boolean Loop |
| --- |
| 파티클의 반복 재생 여부를 설정합니다. |

| float ParticleCount |
| --- |
| 파티클 입자 개수를 설정합니다. |

| float ParticleLifeTime |
| --- |
| 파티클 입자의 지속시간을 설정합니다. |

| float ParticleSize |
| --- |
| 파티클 입자의 크기를 설정합니다. |

| float ParticleSpeed |
| --- |
| 파티클 입자의 속도를 설정합니다. |

| boolean PlayOnEnable |
| --- |
| 파티클 컴포넌트가 Enable일 때, 파티클을 재생할지 여부를 설정합니다. |

| float PlaySpeed |
| --- |
| 파티클 재생 속도를 설정합니다. |

| boolean Prewarm |
| --- |
| 값이 true일 경우 파티클이 처음 재생될 때 이미 재생되고 있었던 것과 같은 상태로 시작합니다. |

| integer RandomSeed |
| --- |
| 파티클이 재생될 때의 생성 위치, 방출 방향, 속도 등을 결정하기 위해 사용하는 랜덤 시드를 설정합니다. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 컴포넌트 활성화 여부를 확인합니다. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| 계층 구조상에서 이 컴포넌트가 Enable 상태인지를 반환합니다. Enable이 true일지라도 엔티티의 Enable이 false라면 false를 반환합니다. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| 이 컴포넌트를 소유한 엔티티입니다. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클을 재생합니다. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| 파티클 재생을 멈춥니다. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| 현재 실행 환경이 클라이언트인지 아닌지의 여부를 반환합니다. |

| boolean IsServer() |
| --- |
| 현재 실행 환경이 서버인지 아닌지의 여부를 반환합니다. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| 파티클의 방출이 종료되었을 때 BaseParticleComponent에서 발생하는 이벤트입니다. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| 파티클 입자 방출이 시작될 때 발생하는 이벤트입니다. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| Loop 프로퍼티가 활성화 된 경우, 파티클의 방출 주기가 돌아와서 방출을 반복할 때 발생하는 이벤트입니다. |

Update 2025-08-27 PM 04:56

