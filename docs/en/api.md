# AIChaseComponent

AI that allows monsters to track players. StateComponent is added automatically when it is not present.

# Properties

| float DetectionRange |
| --- |
| Range of trace detection. If the target moves beyond that range, the trace will be paused. If the target moves back within range, it will resume the trace. |

| boolean IsChaseNearPlayer |
| --- |
| If the value is true, it automatically tracks the nearest player within the DetectionRange property value. If there is a target specified by the TargetEntityRef property or the SetTarget(Entity) function, it will track the specified target instead of the player. |

| [EntityRef](https://mod-developers.nexon.com/apiReference/Misc/EntityRef) TargetEntityRef ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Designates the Entity to be tracked. |

##### inherited from AIComponent:

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether to support the legacy system. Previous systems must use the ExclusiveExecutionWhenRunning property in order to set the node as Running. The legacy system is no longer supported and will be deleted at a later date. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Outputs execution information into a log when the value is true while BehaviorTree runs. Operates in Maker mode only. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Update permission. Executed on the specified location (Client or Server). |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetCurrentTarget() |
| --- |
| Returns the entity being tracked. |

| void SetTarget([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity) |
| --- |
| Sets it to chase targetEntity. The IsChaseNearPlayer property is automatically inactive when calling the function |

##### inherited from AIComponent:

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Creates an Action node. Calls function that had been passed to onBehaveFunction when the node is executed. Parameter of onBehaveFunction is delta, which means time per frame. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| Creates an Action node based on BTNodeType. The nodeType is the type name of BTNodeType. If the onBehaveFunction isn't nil, the function that was passed to the onBehaveFunction will be called instead of functions OnInit() and OnBehave() of the BTNodeType. The parameter of the onBehaveFunction is delta, which means time per frame. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| Sets the node as the top-level node. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of a monster chasing the target that attacked it instead of blindly chasing nearby enemies. You can add and use the component written in the example code to the Chase monster.

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
void OnUpdate ( number delta )
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

Uses BehaviorTree to give AI to an Entity.

# Properties

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether to support the legacy system. Previous systems must use the ExclusiveExecutionWhenRunning property in order to set the node as Running. The legacy system is no longer supported and will be deleted at a later date. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Outputs execution information into a log when the value is true while BehaviorTree runs. Operates in Maker mode only. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Update permission. Executed on the specified location (Client or Server). |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Creates an Action node. Calls function that had been passed to onBehaveFunction when the node is executed. Parameter of onBehaveFunction is delta, which means time per frame. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| Creates an Action node based on BTNodeType. The nodeType is the type name of BTNodeType. If the onBehaveFunction isn't nil, the function that was passed to the onBehaveFunction will be called instead of functions OnInit() and OnBehave() of the BTNodeType. The parameter of the onBehaveFunction is delta, which means time per frame. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| Sets the node as the top-level node. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of a monster sleeping and telling the player to back off if the player approaches. AIComponent is extended and used. Add the created script component to the static monster to use it.

```
Property:
[Sync]
number DetectDistance = 4

Function:
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
- [Creating AI with Behavior Tree](/docs?postId=562)

Update 2025-08-27 PM 04:56


# AIWanderComponent

AI that makes monsters wander around. Added automatically when StateComponent is not present.

# Properties

##### inherited from AIComponent:

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether to support the legacy system. Previous systems must use the ExclusiveExecutionWhenRunning property in order to set the node as Running. The legacy system is no longer supported and will be deleted at a later date. |

| boolean LogEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Outputs execution information into a log when the value is true while BehaviorTree runs. Operates in Maker mode only. |

| [UpdateAuthorityType](https://mod-developers.nexon.com/apiReference/Enums/UpdateAuthorityType) UpdateAuthority ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Update permission. Executed on the specified location (Client or Server). |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from AIComponent:

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateLeafNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName, func<float> -> BehaviourTreeStatus onBehaveFunction) |
| --- |
| Creates an Action node. Calls function that had been passed to onBehaveFunction when the node is executed. Parameter of onBehaveFunction is delta, which means time per frame. |

| [BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) CreateNode([string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeType, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nodeName = nil, func<float> -> BehaviourTreeStatus onBehaveFunction = nil) |
| --- |
| Creates an Action node based on BTNodeType. The nodeType is the type name of BTNodeType. If the onBehaveFunction isn't nil, the function that was passed to the onBehaveFunction will be called instead of functions OnInit() and OnBehave() of the BTNodeType. The parameter of the onBehaveFunction is delta, which means time per frame. |

| void SetRootNode([BTNode](https://mod-developers.nexon.com/apiReference/Misc/BTNode) node) |
| --- |
| Sets the node as the top-level node. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# SeeAlso

- [StateComponent](https://mod-developers.nexon.com/apiReference/Components/StateComponent)

Update 2025-08-27 PM 04:56


# AreaParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the ability to set and control the particle's spawn range, including snow, rain, or cloud.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the location of the center point of the creation scope relative to the Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the particle spawning range. |

| [AreaParticleType](https://mod-developers.nexon.com/apiReference/Enums/AreaParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the type of particle to be created. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles repeatedly. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the number of particles. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle duration. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle size. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the play speed of particles. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined by the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| Play stopped particles. |

| void Stop() |
| --- |
| Stop playing particles. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of playing back and pausing the particle and adjusting AreaSize.

```
Property:
[Sync]
AreaParticleComponent particleComponent = nil

Function:
[client only]
void OnBeginPlay
{
	self.particleComponent = self.Entity.AreaParticleComponent
	 
	self.particleComponent.AreaSize.x = _UtilLogic:RandomDouble() * 2
	self.particleComponent.ParticleCount = 2
}

Event Handler:
[service] [InputService]
{
	----------------- Native Emitter Info ------------------
	-- Emitter: InputService
	-- Space: Client
	--------------------------------------------------------
	 
	-- Parameters
	local key = event.key
	--------------------------------------------------------
	if key == KeyboardKey.Q then
	self.particleComponent:Stop()
	elseif key == KeyboardKey.E then
	self.particleComponent:Play()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)
- [Using Particles](/docs?postId=1036)
- [Utilizing Particles](/docs?postId=764)

Update 2025-12-02 PM 01:55


# AttackComponent

Provide an interface that can reproduce the attack feature linking with HitComponent.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) Attack([Shape](https://mod-developers.nexon.com/apiReference/Misc/Shape) shape, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| Calls the OnHit(Entity, Integer, boolean, string, int32) function of HitComponent within the shape area and triggers HitEvent. Returns all the HitComponents that are considered as attack targets.<br>attackInfo is custom data that can be used as needed by the creator when implementing the attack directly. When utilizing this, the function must be overridden. |

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) Attack([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) size, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) offset, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| An attack function that can define the rectangular area. Size refers to the size of the rectangle. Offset is the entity rectangle's center point position. |

| void AttackFast([Shape](https://mod-developers.nexon.com/apiReference/Misc/Shape) shape, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| Performs an attack, no return value. You can improve the World's function by reducing unnecessary table entity creation. |

| [table<Component>](https://mod-developers.nexon.com/apiReference/Lua/table) AttackFrom([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) size, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) collisionGroup = nil) |
| --- |
| An attack function that defines the rectangular area. Size refers to the size of the rectangle. Position is the entity rectangle's center point position based on the world coordinates. |

| boolean CalcCritical([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Determines whether to launch a critical attack or not. A critical attack is not launched by default behavior, as it always returns false. |

| integer CalcDamage([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Determines the damage value. Default behavior always returns 1. |

| float GetCriticalDamageRate() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| A critical attack determines how many times the default damage will be multiplied by. The default behavior always returns 2. |

| int32 GetDisplayHitCount([string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Calculates how many hits to be displayed from one attack. Default behavior always returns 1. |

| boolean IsAttackTarget([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Determines whether the defender is a valid target for attack. When it returns false, it's removed from the Attack(), AttackFrom(), or AttackFast() functions. The default behavior returns false if the current state of the defender's StateComponent is 'DEAD' and both the defender and the opponent are players, while the opponent's PlayerComponent has a PVPMode property value as false. |

| void OnAttack([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) defender) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Entity attacks trigger this function. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [AttackEvent](https://mod-developers.nexon.com/apiReference/Events/AttackEvent) |
| --- |
| This event occurs when an entity attacks. Occurs in AttackComponent. |

# Examples

By extending AttackComponent, you can implement a custom attack method. Attack monsters entering the front (1, 1) square area.

```
Property:
Function:
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
 
-- Defines how damage is calculated
override int CalcDamage(Entity attacker,Entity defender,string attackInfo)
{
	return 50
}
 
-- Defines how critical damage is calculated.
override boolean CalcCritical(Entity attacker,Entity defender,string attackInfo)
{
	return _UtilLogic:RandomDouble() < 0.3
}
 
-- Defines the probability of critical damage occurring.
override number GetCriticalDamageRate()
{
	return 2
}
 
Entity Event Handler:
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
- [Attack and Hit](/docs?postId=206)

Update 2025-10-28 PM 02:21


# AvatarBodyActionSelectorComponent

This Component selects and applies the avatar's body Action.

# Properties

| [MapleAvatarBodyActionState](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarBodyActionState) ActionState |
| --- |
| Current state is Action. |

| [ReadOnlyDictionary<MapleAvatarBodyActionState, ReadOnlyList<MapleAvatarActionStateElement>>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyDictionary-2) StateActionDic ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Information value about the state. This will be improved later. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ActionStateChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ActionStateChangedEvent) |
| --- |
| This event occurs when the action state changes. |

Update 2025-08-27 PM 04:56


# AvatarFaceActionSelectorComponent

This Component selects and applies the avatar's face Action.

# Properties

| [MapleAvatarFaceActionState](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarFaceActionState) ActionState |
| --- |
| Current state is Action. |

| float BlinkInterval |
| --- |
| Interval between blinks. The unit of measure is the second. |

| [ReadOnlyDictionary<MapleAvatarFaceActionState, ReadOnlyList<MapleAvatarActionStateElement>>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyDictionary-2) StateActionDic ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Information value about the state. This will be improved later. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# AvatarGUIRendererComponent

Renders an avatar-format Entity in the UI.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Changes the entire avatar color to the specified color value. |

| boolean FlipX |
| --- |
| Determines whether to invert based on the X axis. |

| boolean FlipY |
| --- |
| Determines whether to invert based on the Y axis. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| float PlayRate |
| --- |
| You can specify the playback speed of avatar animations. It supports a value of 0 or more. The higher the number, the faster it is. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveAvatar |
| --- |
| Defines how the image's proportions, size, and pivot are saved. |

| boolean RaycastTarget |
| --- |
| Becomes the subject of screen touch or mouse clicks if the value is set to true. The UI hidden behind the avatar will not receive screen touch and mouse click inputs. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetAvatarRootEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets the AvatarRoot Entity, the top-level Entity composing the avatar. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetBodyEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets the body Entity composing the avatar. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetFaceEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets the face Entity composing the avatar. |

| void PlayEmotion([EmotionalType](https://mod-developers.nexon.com/apiReference/Enums/EmotionalType) emotionalType, float duration) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Play emotion. |

| void SetAvatarPartColor([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, float r, float g, float b, float a) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the color values of the parts corresponding to the avatar's category. Uses RGB values. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

#### PlayEmotion

Press the F1 through F5 keys to change emotes.

```
Method:
[client only]
void ChangeEmotion(number emotionNumber)
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
[client only] [InputService]
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
- [Representing Avatars in the UI](/docs?postId=953)

Update 2025-11-24 AM 11:42


# AvatarRendererComponent

Renders an avatar-format Entity.

# Properties

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Doesn't perform automatic replace when designating the Map Layer name into SortingLayer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| int32 OrderInLayer |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Control the speed of the avatar's animation. |

| boolean ShowDefaultWeaponEffects ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets up a weapon's default effect animation and basic sounds. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Avatar's SortingLayer. When two or more Entities overlap, the display priority is determined according to the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetAvatarRootEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets the top-level Entity making up the avatar. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetBodyEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets an Entity corresponding to the avatar's body. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) GetFaceEntity() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Gets an Entity corresponding to the avatar's face. |

| void PlayEmotion([EmotionalType](https://mod-developers.nexon.com/apiReference/Enums/EmotionalType) emotionalType, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Plays emotional expressions. |

| void SetAlpha(float alpha, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Sets the avatar's Alpha value. |

| void SetAvatarPartColor([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, float r, float g, float b, float a, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Sets the color values of the parts corresponding to the avatar's category. Uses RGB values. |

| void SetColor(float r, float g, float b, float a, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Sets the avatar's color value. Uses RGB values. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

Update 2025-08-27 PM 04:56


# AvatarStateAnimationComponent

Specifies an animation to be played according to the Avatar's state changes.

# Properties

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether or not to support the legacy system. The previous system is no longer supported and will be deleted at a later date. |

| [SyncDictionary<string, AvatarBodyActionElement>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) StateToAvatarBodyActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| This is a table in which state names and AvatarBodyActionState are mapped. Used when IsLegacy's value is false. |

##### inherited from StateAnimationComponent:

| [SyncDictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The table the animation's name and AnimationClip are mapped to. Used when IsLegacy's value is true. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Function which handles StateChangeEvent upon receiving. AnimationClipEvent is raised to play AnimationClip which is basically mapped to a State. |

##### inherited from StateAnimationComponent:

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Function which handles StateChangeEvent upon receiving. AnimationClipEvent is raised to play AnimationClip which is basically mapped to a State. |

| void RemoveActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key) |
| --- |
| Removes the element corresponding to key in StateToAvatarBodyActionSheet. Removes the element from ActionSheet if IsLegacy's value is true. |

| void SetActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) animationClipRuid) |
| --- |
| Adds elements to StateToAvatarBodyActionSheet. AvatarBodyActionStateName's property value of the AvatarBodyActionElement entity added as an element is animationClipRuid, and PlayerRate's property value is 1. Adds an element to ActionSheet when when IsLegacy's value is true. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) StateStringToAnimationKey([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Returns the name of mapped Animation to the State. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [BodyActionStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/BodyActionStateChangeEvent) |
| --- |
| This event occurs when the BodyAction state changes. |

Update 2025-08-27 PM 04:56


# BackgroundComponent

Manages the background of the map. Solid color backgrounds, MapleStory backgrounds, and web image backgrounds are available.

# Properties

| [SyncDictionary<string, BackgroundPieceDataElement>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) BackgroundPieces ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| This is the Background Piece list for the background. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| A color value at the bottom of the gradient background. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) MiddleBottomColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| A color value at the bottom middle of the gradient background. |

| float MiddleBottomRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The reference position of the color at the bottom middle of the gradient background. You can set a value between 0.01 and MiddleTopRatio - 0.01. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) MiddleTopColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| A color value at the top middle of the gradient background. |

| float MiddleTopRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The reference position of the color at the top middle of the gradient background. You can set a value between MiddleBottomRatio + 0.01 and 0.99. |

| float ScrollRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| You can adjust the movement speed of the background when the Type is Template. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) SolidColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The color value with SolidColor Type. To change the background, use the ChangeBackgroundBySolidColor function. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) TemplateRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| RUID with Template Type. To change the background, use the ChangeBackgroundByTemplateRUID function. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| A color value at the top of the gradient background. |

| [BackgroundType](https://mod-developers.nexon.com/apiReference/Enums/BackgroundType) Type ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The type of background. The background can be changed by using the ChangeBackgroundBy() functions. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) WebUrl ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The URL of the web image background when the Type is Web. To change the background, use the ChangeBackgroundByWebUrl function. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeBackgroundByGradient([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) top, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) middleTop, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) middleBottom, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) bottom, float middleTopRatio, float middleBottomRatio) |
| --- |
| Changes the background to gradient colors. |

| void ChangeBackgroundBySolidColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) value) |
| --- |
| Changes the background to a solid color. |

| void ChangeBackgroundByTemplateRUID([string](https://mod-developers.nexon.com/apiReference/Lua/string) value) |
| --- |
| Changes the background according to the specified Template background RUID. |

| void ChangeBackgroundByWebUrl([string](https://mod-developers.nexon.com/apiReference/Lua/string) value) |
| --- |
| Changes the background to a web image. |

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| void SetBackgroundPieceColor([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the color by searching for the Background Piece that corresponds to pieceName. |

| void SetBackgroundPieceEnable([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, boolean enable) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets activation status by searching for the Background Piece that corresponds to pieceName. |

| void SetBackgroundPiecePosition([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets location by searching for the Background Piece that corresponds to pieceName. |

| void SetBackgroundPieceRatio([string](https://mod-developers.nexon.com/apiReference/Lua/string) pieceName, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ratio) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the background piece ratio by searching for the Background Piece that corresponds to pieceName. This ratio determines how much the Background Piece moves when the camera moves. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

The following example shows various methods of changing the background.

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

- [Set Background](/docs?postId=768)

Update 2025-12-02 PM 01:55


# BaseParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)![custom](https://img.shields.io/static/v1?label=&amp;message=Abstract&amp;color=darkkhaki)

This is the parent component of ParticleComponents that make a particle effect.

# Properties

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles repeatedly. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the number of particles. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle duration. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle size. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the play speed of particles. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined by the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void Play() |
| --- |
| Play stopped particles. |

| void Stop() |
| --- |
| Stop playing particles. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

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

Provides setting and control of basic particles.

# Properties

| [BasicParticleType](https://mod-developers.nexon.com/apiReference/Enums/BasicParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the type of particle to be created. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles repeatedly. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the number of particles. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle duration. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle size. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the play speed of particles. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined by the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| Play stopped particles. |

| void Stop() |
| --- |
| Stop playing particles. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of playing back and pausing the particle and controlling the particle property.

```
Property:
[Sync]
BasicParticleComponent particleComponent = nil

Method:
[client only]
void OnBeginPlay
{
	self.particleComponent = self.Entity.BasicParticleComponent
	 
	self.particleComponent.PlaySpeed = _UtilLogic:RandomDouble() * 2
	self.particleComponent.ParticleCount = 2
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
	self.particleComponent:Stop()
	elseif key == KeyboardKey.E then
	self.particleComponent:Play()
	end
}
```

# SeeAlso

- [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey)
- [UtilLogic](https://mod-developers.nexon.com/apiReference/Logics/UtilLogic)
- [InputService](https://mod-developers.nexon.com/apiReference/Services/InputService)
- [ParticleService](https://mod-developers.nexon.com/apiReference/Services/ParticleService)
- [Using Particles](/docs?postId=1036)
- [Utilizing Particles](/docs?postId=764)

Update 2025-12-02 PM 01:55


# ButtonComponent

Provides UI button functions.

# Properties

| TransitionColorSet Colors ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the button's color according to its state. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| TransitionRUIDSet ImageRUIDs ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the image of the button according to its state. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it has been placed in the world. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) KeyCode |
| --- |
| When pressing the button, it operates as if the specified keyCode was pressed. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| [TransitionType](https://mod-developers.nexon.com/apiReference/Enums/TransitionType) Transition |
| --- |
| This is a type for the button's state transition. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ButtonClickEditorEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonClickEditorEvent) |
| --- |
| The editor event occurs when clicking the Button. |

| [ButtonClickEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonClickEvent) |
| --- |
| The event occurs when clicking a Button. |

| [ButtonPressedEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonPressedEvent) |
| --- |
| This event occurs when the button's state becomes Pressed. |

| [ButtonStateChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonStateChangeEditorEvent) |
| --- |
| The editor event occurs when changing the Button's state. |

| [ButtonStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/ButtonStateChangeEvent) |
| --- |
| This event occurs when changing the Button's state. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

Here's an example of changing the button image or the color of the image while pressing the button. You can test by adding this Component to an Entity with ButtonComponent.

```
Property:
[None]
boolean IsButtonDown = false
[None]
number RedVal = 0
[None]
number TimerID = 0
[None]
string MonsterRUID = "96e955c1bf27415e84f96deea200a8f1"
[None]
string OriginalRUID = ""
  
  
Method:
[client only]
void OnBeginPlay()
{
	self.OriginalRUID = self.Entity.SpriteGUIRendererComponent.ImageRUID
}
  
void CancelHoldButton()
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
- [Basic UI Components](/docs?postId=744)

Update 2025-12-02 PM 01:55


# CameraComponent

Adds a camera function that faces an Entity. Use CameraService to switch between cameras.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CameraOffset |
| --- |
| Sets the position of the camera based on World coordinates. |

| boolean ConfineCameraArea |
| --- |
| Sets the camera to cover only the foothold area of the map. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Damping |
| --- |
| Determines how quickly the camera responds when a subject enters the SoftZone while it is tracking the subject. The smaller the value, the more responsive it is. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DeadZone |
| --- |
| Sets the DeadZone area, the section of the frame where the camera holds its target. |

| float DutchAngle |
| --- |
| Sets the rotation value of the camera. |

| boolean IsAllowZoomInOut |
| --- |
| The active camera is checked to see whether zoom operation is allowed. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LeftBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The bottom-left value of the camera-restricted area. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RightTop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The top-right value of the camera-restricted area. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ScreenOffset |
| --- |
| The ratio value of the full screen based on the target. Values can be between 0 and 1; 0.5 will center the camera. Available when ConfineCameraArea is false. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SoftZone |
| --- |
| Sets the SoftZone area. If target enters the frame area, the camera will shift direction to DeadZone. |

| boolean UseCustomBound ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shows whether you will directly define and use the camera-restricted area. You can define it using the LeftBottom and RightTop properties. When this is False, you will use the map area of the map where you belong. If the map area is the default, the corrected area based on that map area will be used as the camera-restricted area. |

| float ZoomRatio |
| --- |
| Sets the zoom ratio. The value must be set to ZoomRatioMin or greater and ZoomRationmax or less in percentage. |

| float ZoomRatioMax |
| --- |
| Sets the maximum value for the camera's zoom ratio. The value cannot exceed 500. |

| float ZoomRatioMin |
| --- |
| Sets the minimum value for the camera's zoom ratio. The value must be greater than 30. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| Vector2, Vector2 GetBound() |
| --- |
| Gets the camera-restricted area composed of LeftBottom and RightTop. |

| void SetZoomTo(float percent, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| The camera will zoom in for the given amount of time (seconds). |

| void ShakeCamera(float intensity, float duration, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Vibrates the camera for a given amount of time (seconds). |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This example sets the zoom factor to 300% after 4 seconds.

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
- [Controlling the Camera with CameraService](/docs?postId=118)

Update 2025-08-27 PM 04:56


# CanvasGroupComponent

Used to control an entire group of UI elements in one place.

# Properties

| boolean BlocksRaycasts |
| --- |
| When it's set as true, the child UI can receive a screen touch or click input. The UI hidden behind will not receive the input. |

| float GroupAlpha |
| --- |
| Sets the child UI's transparency to a value between 0 and 1. |

| boolean Interactable |
| --- |
| Sets whether the child UI is interactive. If it's set as false, it will not operate even if it receives the user input. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# SeeAlso

- [Basic UI Components](/docs?postId=744)
- [Creating UI](/docs?postId=64)

Update 2025-10-28 PM 02:21


# ChatBalloonComponent

Displays a chat balloon and provides a feature to set the related function.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the Message property. |

| boolean ArrowChatEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to provide the chat balloon a tail image or not. |

| boolean AutoShowEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to display chat balloons automatically or not. If the value is true, chat balloons will be visible during ShowDuration and invisible during HideDuration. This action will be on repeat. If a player's ChatModeEnabled value is true, automatic chat balloon feature will not operate. |

| float BalloonScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the size of the chat balloon. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ChatBalloonRUID |
| --- |
| Sets the type of the chat balloon. |

| boolean ChatModeEnabled ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Decides whether to link with chat. If true, the text entered on the chat window gets displayed as a chat balloon and only works in the Player. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the text color of the chat balloon. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) FontOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Adjusts the position of the text within the chat balloon. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the font size in the chat balloon. |

| float HideDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Decides the amount of time that the chat balloon remains invisible when AutoShowEnabled's value is set to true. After the set time has expired, a chat balloon appears. Defined in seconds. |

| boolean IsRichText ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to use rich text. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Message ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Decides the content that will be displayed in the chat balloon. A chat balloon will not appear when it's nil or an empty string. |

| float Offset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Position offset of the chat balloon. You can adjust the position of the chat balloon up or down. |

| float ShowDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| You can set the duration of the chat balloon. A chat balloon appears and disappears after Delay. Defined in seconds. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ChatBalloonEvent](https://mod-developers.nexon.com/apiReference/Events/ChatBalloonEvent) |
| --- |
| This event occurs when a chat balloon appears. |

# Examples

The following is an example of making an NPC that speaks lines at regular intervals. ChatBalloon replaces the dialogue according to the output interval.

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
- [Making Chat Balloons](/docs?postId=119)

Update 2025-10-28 PM 02:21


# ChatComponent

The component supports the chat function for players to communicate with each other.

# Properties

| float ChatEmotionDuration |
| --- |
| Sets the avatar emote duration. |

| boolean EnableVoiceChat ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to show and use the voice chat button. |

| boolean Expand |
| --- |
| The function expands the chat window. |

| boolean HideWorldChatButton ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| This feature allows you to hide the World chat button. |

| boolean MessageAlignBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| This feature allows you to arrange the World chat message to the bottom. |

| boolean UseChatBalloon |
| --- |
| The function for players to express a chat message as a chat balloon. |

| boolean UseChatEmotion |
| --- |
| The function to use avatar emotes in chat. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ChatEvent](https://mod-developers.nexon.com/apiReference/Events/ChatEvent) |
| --- |
| The event occurs when entering a dialog. |

# Examples

This example adjusts ShowDuration for ChatEmotionDuration and ChatBalloonComponent based on EmotionalType length, and verifies EmotionalType presence in the message via ChatEvent.

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

Specifies the area for climbing action.

# Properties

| boolean AllowHorizontalMove |
| --- |
| Sets whether or not to move freely. If true, both X and Y axes are movable. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the center point position of the collider rectangle based on Entity. Setting IsUseDefaultObjectBoxSize to false can apply BoxOffset. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the collider rectangle. Setsting IsUseDefaultObjectBoxSize to false can apply BoxSize. |

| [ClimbableType](https://mod-developers.nexon.com/apiReference/Enums/ClimbableType) ClimbableAnimationType |
| --- |
| Determines the animation type when climbing. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| A group of collisions in the climbing area. |

| boolean IsUseDefaultObjectBoxSize |
| --- |
| Sets whether to fit the climbing area settings to the sprite size. Operates only in Maker Editing mode. If true, automatically sets to fit the Boxsize and BoxOffset to the sprite size. If false, you can set the Boxsize and BoxOffset arbitrarily. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SpeedFactor |
| --- |
| When climbing, this value is multiplied to the speed traveled in the X or Y direction. The greater the value, the faster the movement. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example where `BoxSize` and `SpeedFactor` of `ClimbableComponent` change based on the name of the colliding entity when `TriggerEnterEvent` occurs.

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
- [Using Ladders and Rope](/docs?postId=809)

Update 2025-08-27 PM 04:56


# ClimbableSpriteRendererComponent

Sets up a Climbable Sprite and display the image on the world.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ClipName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Colors the sprite. In case of basic white, the original color is displayed. |

| [SpriteDrawMode](https://mod-developers.nexon.com/apiReference/Enums/SpriteDrawMode) DrawMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the Sprite's draw mode. Tiled can be used. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether the sprite is inverted relative to the X axis. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether the sprite is inverted relative to the Y axis. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Doesn't perform automatic replace when designating the Map Layer name into SortingLayer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material that the renderer will use. |

| boolean NeedGizmo ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays the Sprite's Gizmo. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same Layer. A greater number indicates higher priority. |

| RenderSettingType RenderSetting ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Sprite RUID for the ladder body. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUIDHead ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the Sprite RUID for the ladder head. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUIDTail ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the Sprite RUID for the ladder tail. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TiledSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| You can set the vertical height of the ladder by changing the Y value. Changing the X value is not supported. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| void ResetColliderBox() |
| --- |
| Initializes the collision area to fit the ladder's sprite size. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [EmbededSpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeFrameEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerChangeFrameEvent. |

| [EmbededSpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeStateEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerChangeStateEvent. |

| [EmbededSpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerEndEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerEndEvent. |

| [EmbededSpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerStartEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerStartEvent. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

| [SpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeFrameEvent) |
| --- |
| This event is raised when changing the sprite animation's frame. |

| [SpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeStateEvent) |
| --- |
| This event is deprecated. |

| [SpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndEvent) |
| --- |
| This event is raised when finishing the sprite animation's playback. |

| [SpriteAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndFrameEvent) |
| --- |
| This event is raised when playing the last frame of sprite animation. |

| [SpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartEvent) |
| --- |
| This event is raised when starting the sprite animation's playback. |

| [SpriteAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartFrameEvent) |
| --- |
| This event is raised when playing the first frame of sprite animation. |

# Examples

This is an example of changing the `TiledSize` of `ClimbableRendererComponent` and adjusting the collider box's size based on the sprite size when the `TriggerEnterEvnet` occurs, depending on the name of the collided entity.

```
Function:
[server]
void ResizeRenderer (Vector2 tiledSize)
{
	local entity = _EntityService:GetEntityByPath(EntityPath)
	entity.ClimbableSpriteRendererComponent.TiledSize = tiledSize
	entity.ClimbableSpriteRendererComponent:ResetColliderBox()
}

Entity Event Handler:
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
- [Using Ladders and Rope](/docs?postId=809)

Update 2025-08-27 PM 04:56


# Component

The parent component of all Components. Provides basic functions of Components.

# Properties

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# CostumeManagerComponent

Manages information such as clothes and weapons equipped by the player.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomBodyEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Skin |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCapeEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Cape |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCapEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Hat |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomCoatEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Coat |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEarAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Ear Accessory |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEarEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Ears |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomEyeAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Eye Accessory |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomFaceAccessoryEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Face Accessory |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomFaceEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Face |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomGloveEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Gloves |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomHairEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Hair Style |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomLongcoatEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Long Coat |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomOneHandedWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 1H Weapon |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomPantsEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Pants |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomShoesEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shoes |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomSubWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sub-weapon |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CustomTwoHandedWeaponEquip ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| 2H Weapon |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) DefaultEquipUserId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Custom equipment will be applied after duplicating the target user's equipment. Even the users not connecting can designate it, and it won't be applied if the target user's equipment is changed after the designation. |

| [ReadOnlyList<MapleAvatarItemData>](https://mod-developers.nexon.com/apiReference/Misc/ReadOnlyList-1) EquippedItems ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Information about equipped items. It cannot be modified in the script. |

| boolean UseCustomEquipOnly ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| It only uses costume specified in the script, not the default costume of the user. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetEquip([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the itemRUID currently being equipped. Access is available using property. |

| void SetEquip([MapleAvatarItemCategory](https://mod-developers.nexon.com/apiReference/Enums/MapleAvatarItemCategory) category, [string](https://mod-developers.nexon.com/apiReference/Lua/string) itemRUID) |
| --- |
| Equips the item corresponding to the entered itemRUID. Enter an empty string to unlock the device. Using properties can change it. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [DestroyMapleCostumeEvent](https://mod-developers.nexon.com/apiReference/Events/DestroyMapleCostumeEvent) |
| --- |
| Event that occurs when CostumeManagerComponent is deleted. |

| [InitMapleCostumeEvent](https://mod-developers.nexon.com/apiReference/Events/InitMapleCostumeEvent) |
| --- |
| The event occurs when changing the wearing state of equipment. |

# Examples

This example shows how a player’s hair changes and outputs the RUID when the player touches another entity.

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

A custom foothold Component. You can set to have a physical effect on an Entity on the foothold.

# Properties

| [SyncList<List<Vector2>>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) edgeLists ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Indicates a foothold. A foothold consists of several connecting points. |

| float FootholdDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| A friction force is applied to an Entity containing a RigidbodyComponent on the foothold. The larger the value, the faster it decelerates. |

| float FootholdForce ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| A force is applied to an Entity containing a RigidbodyComponent on the foothold. Move to the right if the value is positive, and to the left if negative. |

| float FootholdWalkSpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The coefficient multiplied by the movement speed when an Entity with the RigidbodyComponent is on the foothold. The higher the value is, the faster the movement speed. |

| boolean IsBlockVerticalLine ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether the entity is blocked by the vertical platforms of the tile map. Applies only to entities with a RigidbodyComponent. |

| boolean IsDynamicFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Can move the foothold or change the shape while playing if the value is true. Please be aware that changing the foothold's location or shape frequently may negatively affect your World's function. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| If true, this may collide with a Dynamic rigid body (PhysicRigidbody) using the Physics feature. |

| [RigidbodyMovementOptionType](https://mod-developers.nexon.com/apiReference/Enums/RigidbodyMovementOptionType) RigidbodyMovementOption ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets options to apply to the RigidbodyComponent above the foothold. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

You can change the foothold's shape by modifying `edgeList`. When you change the foothold, please note the following:

1. Changes will be applied only when `IsDynamicFoothold` is true.
2. You must create a new list and overwrite it to apply the changes. Please refer to the example code below.

An example of changing part of the foothold.

```
local target = _EntityService:GetEntityByPath(EntityPath)
 
-- If you modify the element's value, the changes will not be applied to the foothold.
-- target.CustomFootholdComponent.edgeLists[1][1] = Vector2(3, 0)
 
-- You must create a new list and overwrite it.
local list = {Vector2(3,0), Vector2(2,3)}
target.CustomFootholdComponent.edgeLists[1] = list
```

An example of replacing the entire foothold.

```
local listlist = {{Vector2(0,2), Vector2(3,2)},{Vector2(4,5), Vector2(3,8)}}
local target = _EntityService:GetEntityByPath(EntityPath)
target.CustomFootholdComponent.edgeLists = listlist
```

# SeeAlso

- [EntityService](https://mod-developers.nexon.com/apiReference/Services/EntityService)
- [Making a Moving Foothold](/docs?postId=579)
- [Making Footholds](/docs?postId=71)

Update 2025-12-02 PM 01:55


# DamageSkinComponent

Configure damage skins in this component to visually express damage. Specify the format in the attacker Entity's DamageSkinSettingComponent.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# DamageSkinSettingComponent

Specifies the damage skin type for attacks. The attacked entity must have DamageSkinSpawnerComponent.

# Properties

| float Alpha ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the Alpha value of the damage skin. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) DamageSkinId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the damage skin's type. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DamageSkinScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the damage skin's size. |

| float DelayPerAttack ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the damage skin's delay time in seconds. Please refer to the GetDisplayHitCount(attackInfo) function of AttackComponent. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the playback speed of the damage skin. |

| [DamageSkinTweenType](https://mod-developers.nexon.com/apiReference/Enums/DamageSkinTweenType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the damage skin's movement format. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# DamageSkinSpawnerComponent

Outputs the damage skin when a hit event occurs on the Entity.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) DamageSkinOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the location for the damage skin to be output. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example where damage is output to the top center based on the ColliderOffset and BoxSize of the current HitComponent.

```
Function:
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

Follows the parent Entity’s direction. The current Entity’s orientation changes with a ChangedLookAtEvent in the parent Entity.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

You can execute the example below after adding the object-1 model to the workspace. Create an object facing the same direction as the player by adding the Component below to DefaultPlayer.

```
Function:
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

You can create and delete DistanceJoint. Keeps the distance between connected rigid bodies constant.

# Properties

| [SyncList<DistanceJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Set Joint information. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, float length) |
| --- |
| Add DistanceJoint. Return index on success, -1 on failure |

| void DestroyJoint(int32 index) |
| --- |
| Removes the Joint whose sequence number corresponds to the index. |

| int32 GetJointsCount() |
| --- |
| Return the number of joints. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| Set the CollideConnected value of the Joint whose sequence number corresponds to index. |

| void SetLength(int32 index, float length) |
| --- |
| Set the Length value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| Set the LocalAnchorA value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Set the LocalAnchorB value of the Joint whose sequence number corresponds to index. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

#### AddJoint

This is an example of pressing R, adding Joint, and catching the falling ball.

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

Entity Event Handler:
[service] [InputService]
HandleKeyDownEvent ( KeyDownEvent event )
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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-10-28 PM 02:21


# FootholdComponent

Manages all the footholds in the map. Footholds interact with RigidbodyComponent.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetFoothold(int32 footholdId) |
| --- |
| Returns Foothold corresponding to footholdId. |

| [table<Foothold>](https://mod-developers.nexon.com/apiReference/Lua/table) GetFootholdAll() |
| --- |
| Returns all Footholds on the map. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetNearestFootholdByPoint([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, float distance) |
| --- |
| Finds and returns the closest Foothold within a certain distance from the specified point. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) Raycast([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float distance) |
| --- |
| Fires a moving ray over a distance in the direction from the point location and then finds and returns the first Foothold colliding with this ray. |

| [table<Foothold>](https://mod-developers.nexon.com/apiReference/Lua/table) RaycastAll([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) point, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float distance) |
| --- |
| Fires a moving ray over a distance in the direction from the point location and then finds and returns all the first Footholds colliding with this ray. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

Touching the screen randomly selects a number and matches `GetNearestFootholdByPoint`, `Raycast`, and `RaycastAll`. This example shows how to destroy the dynamic foothold detected by these three functions.

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
- [Making a Moving Foothold](/docs?postId=579)
- [Making Footholds](/docs?postId=71)

Update 2025-10-28 PM 02:21


# GridViewComponent

It is a Component that expresses a standardized UI Entity in the form of a grid. Since only a UI Entity visible on the screen is created and reused, it is an optimized form to express a large amount of Grid.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CellSize |
| --- |
| Fixed size of child UI Entity. |

| int32 FixedCount |
| --- |
| Sets a fixed number of rows or columns. |

| [GridViewFixedType](https://mod-developers.nexon.com/apiReference/Enums/GridViewFixedType) FixedType |
| --- |
| Sets whether to freeze rows or columns. |

| [HorizontalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/HorizontalScrollBarDirection) HorizontalScrollBarDirection |
| --- |
| The orientation of the horizontal scrollbar. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) ItemEntity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The source that will be cloned when creating a child UI Entity. Entity set as ItemEntity is disabled. |

| function<int32, Entity> OnClear ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| A callback when not using a child UI Entity. The indexes of the child UI Entity and UI Entity are passed. |

| function<int32, Entity> OnRefresh ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| A callback when reusing a child UI Entity. The indexes of the child UI Entity and UI Entity are passed. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| Sets the free space on the top, bottom, left, and right of the grid view. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarBackgroundColor |
| --- |
| The background color of the scrollbar. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarBackgroundImageRUID |
| --- |
| The background image of the scrollbar. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarHandleColor |
| --- |
| The handle color of the scrollbar. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarHandleImageRUID |
| --- |
| The handle image of the scrollbar. |

| float ScrollBarThickness |
| --- |
| The thickness of the scrollbar area. |

| [ScrollBarVisibility](https://mod-developers.nexon.com/apiReference/Enums/ScrollBarVisibility) ScrollBarVisible |
| --- |
| Sets whether to display the scrollbar automatically. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Spacing |
| --- |
| The spacing between child UI Entities. |

| int32 TotalCount |
| --- |
| Total number of child UI Entities to be displayed in GridView. |

| boolean UseScroll |
| --- |
| Sets whether to use the scroll function. |

| [VerticalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/VerticalScrollBarDirection) VerticalScrollBarDirection |
| --- |
| The orientation of the vertical scrollbar. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetScrollNormalizedPosition() |
| --- |
| Returns the normalized position of the scroll bar. |

| float GetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| Returns the normalized position of the specified directional scroll bar. |

| void Refresh(boolean resetPos = true, boolean force = false) |
| --- |
| All the child UI Entities of GridView are updated. If restPos is true, initializes the scroll location after the update. |

| void RefreshIndex(int32 index) |
| --- |
| Updates the child UI Entity at a specific index. The indexes of the child UI Entity and the UI Entity are called through the OnRefresh callback. |

| void ResetScrollPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| Moves the position of the scroll bar on the specified axis to the top position. |

| void SetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis, float value) |
| --- |
| Moves the position of the scrollbar on the specified axis to the designated normalized position. The top is 0, the bottom is 1. |

| void SetScrollPositionByItemIndex(int32 index) |
| --- |
| Moves the scroll bar to a position where the child UI Entity at a specific index is visible. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ScrollPositionChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ScrollPositionChangedEvent) |
| --- |
| An event that occurs when the scroll position changes on a scrollable UI Entity. Only occurs if the UI Entity has ScrollLayoutGroupComponent or GridViewComponent. |

# Examples

This is an example of arranging Entities according to a layout.

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

Sets the entity’s collision area and provides an interface for hit reactions when receiving an attack from AttackComponent.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Available on previous systems where IsLegacy is true. Sets the center point position of the collider rectangle based on Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the rectangular collider. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The radius of the circular collider. Valid when ColliderType is Circle. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. Use CollisionGroup. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the center point position of the collider based on Entity. Available on new systems with IsLegacy set as false. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the type of collider. Available on new systems with IsLegacy set as false. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set up a collision group. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Whether or not this Component works with the old system. The new system allows the Collider to be subject to the TransformComponent's rotation and size. You can also use a circle-shaped collider by setting the ColliderType. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The positions of the points that make up the polygonal collider. Valid when ColliderType is Polygon. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| boolean IsHitTarget([string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Determines whether an Entity is subject to AttackComponent's attack. The default behavior is true.<br>attackInfo is custom data set by the creator that was passed from AttackComponent. |

| void OnHit([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) attacker, integer damage, boolean isCritical, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attackInfo, int32 hitCount) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Called when an Entity is hit. The default action is to raise a HitEvent.<br>attacker is the Entity that initiated the attack, attackInfo is the custom data passed from AttackComponent, and hitCount is the number of times the damage split is played. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [HitEvent](https://mod-developers.nexon.com/apiReference/Events/HitEvent) |
| --- |
| This event occurs when an entity is hit. |

# Examples

This is an example of increasing the BoxSize of `HitComponent` in inverse proportion to remaining health whenever a monster is attacked by a player.

```
Property:
[Sync]
number Health = 1000
[None]
number InitialHealth = 0
[None]
Vector2 InitialBoxSize = Vector2(0,0)
  
Function:
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

`HitComponent` can be extended to directly implement both hit detection and actions to take when hit. The following is the IsHitTarget function of the PlayerHit script, which is a built-in script. This makes the player invincible for 1 second when hit.

```
Property:
[None]
number ImmuneCooldown = 1
[None]
number LastHitTime = 0
  
Function:
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
- [Attack and Hit](/docs?postId=206)

Update 2025-08-27 PM 04:56


# HitEffectSpawnerComponent

Output the hit effect when a hit event occurs on the Entity.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# ImageComponent

This is the parent component for UI image printing components like SpriteGUIRendererComponent. It cannot be added directly to the entity.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Sets the default color for the image. |

| boolean DropShadow |
| --- |
| Sets whether to output shadows for the image. |

| float DropShadowAngle |
| --- |
| Sets the angle to output the shadow. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| Sets the shadow color. |

| float DropShadowDistance |
| --- |
| Distance between image and shadow. |

| float FillAmount |
| --- |
| Percentage of the image with Type setting Filled. Use values between 0 and 1. |

| boolean FillCenter |
| --- |
| Sets whether to fill the center of the image area when Type is set to Sliced or Tiled. |

| boolean FillClockWise |
| --- |
| Sets the direction of filling when FillMethod is set to Radial90, Radial180, or Radial360. It fills clockwise if the value is true. |

| [FillMethodType](https://mod-developers.nexon.com/apiReference/Enums/FillMethodType) FillMethod |
| --- |
| Filling method when setting the Filled type. |

| int32 FillOrigin |
| --- |
| You can set the starting point for filling when Type is set to Filled. If FillMethod is Horizontal or Vertical, you can use values from 0 to 1. If FillMethod is Radial90, Radial180, or Radial360, you can use values from 0 to 3. |

| boolean FlipX |
| --- |
| Determines whether to invert based on the X axis of an image. |

| boolean FlipY |
| --- |
| Determines whether to invert based on the Y axis of an image. |

| int32 FrameColumn ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| int32 FrameRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| int32 FrameRow ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ImageRUID |
| --- |
| The image RUID to be displayed on the screen. |

| boolean Outline |
| --- |
| Sets whether to output the image outline. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| Image outline color. |

| float OutlineWidth |
| --- |
| Outline thickness. |

| boolean RaycastTarget |
| --- |
| Becomes the subject of screen touch or mouse clicks if the value is set to true. The UI hidden behind will not receive screen touch and mouse click inputs. |

| [ImageType](https://mod-developers.nexon.com/apiReference/Enums/ImageType) Type |
| --- |
| How to display the image. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void SetNativeSize() |
| --- |
| Resizes the image to its original size. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# InteractionComponent

Component available to interact with an Entity.

# Properties

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) ActionKey ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the usable Key for Interaction. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ActionName ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the name of the Interaction displayed in the chat balloons. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Available on previous systems where IsLegacy is true. Sets the center point position of the collider rectangle based on Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the rectangular collider. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The radius of the circular collider. Valid when ColliderType is Circle. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. Use CollisionGroup. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the center point position of the collider based on Entity. Available on new systems with IsLegacy set as false. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the type of collider. Available on new systems with IsLegacy set as false. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Collider's collision group. |

| float HoldingDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When InteractionType is set to KeyHoldingDuration or KeyUpAfterHoldingDuration, it sets how long the key should be held down in seconds. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) InfoOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| This property is no longer used. |

| [InteractType](https://mod-developers.nexon.com/apiReference/Enums/InteractType) InteractionType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the key input method required for interaction. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether this Component will operate with the previous system. In new system, the Collider is affected by TransformComponent's rotation and size. You can also use a circle-shaped collider by setting the ColliderType. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The positions of the points that make up the polygonal colliding body. Valid when ColliderType is Polygon. |

| boolean ShowActionInfo ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether a chat balloon will appear and show the ActionName and ActionKey when a player approaches. A ChatBalloonComponent will be added automatically during gameplay to display the chat balloon. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void OnEnter() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is deprecated. Use InteractionEnterEvent. |

| void OnInteraction() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is deprecated. Use InteractionEvent. |

| void OnLeave() ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is deprecated. Use InteractionLeaveEvent. |

| void SetOnEnter(func onEnterFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is deprecated. Use InteractionEnterEvent. |

| void SetOnInteraction(func onInteractionFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is deprecated. Use InteractionEvent. |

| void SetOnLeave(func onLeaveFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is deprecated. Use InteractionLeaveEvent. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [InteractionEnterEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEnterEvent) |
| --- |
| The event occurs when entering the interaction area. |

| [InteractionEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionEvent) |
| --- |
| The Event occurs during Interaction. |

| [InteractionLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/InteractionLeaveEvent) |
| --- |
| The Event occurs when leaving the Interaction area. |

# Examples

When Interaction occurs, the InputSpeed of the MovementComponent for the Player Entity changes to 3.

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

The component which manages owned items.

# Properties

| boolean IsInitialized ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Checks whether inventory initialization is complete. InventoryComponent-related functions will work normally after initialization. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [table<Item>](https://mod-developers.nexon.com/apiReference/Lua/table) GetItemList() |
| --- |
| Gets the owned items. |

| [table<Item>](https://mod-developers.nexon.com/apiReference/Lua/table) GetItemsWithType(Type itemType) |
| --- |
| Brings the input type item among the owned ones. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [InventoryItemAddedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemAddedEvent) |
| --- |
| This event occurs when adding an item to the inventory. |

| [InventoryItemInitEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemInitEvent) |
| --- |
| This event occurs when initializing the inventory. InventoryComponent related functions will work normally after the event. |

| [InventoryItemModifiedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemModifiedEvent) |
| --- |
| This event occurs when editing an item of the inventory. |

| [InventoryItemRemovedEvent](https://mod-developers.nexon.com/apiReference/Events/InventoryItemRemovedEvent) |
| --- |
| This event occurs when removing an item of the inventory. |

# Examples

The InventoryComponent fires events when an item is added, removed, or changed in the inventory. Here is an example of when each event was called using the example code for ItemService.

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

When adding and removing items to the inventory using the `_ItemService:CreateItem` and `_ItemService:RemoveItem` functions, the `InventoryItemAddedEvent` and `InventoryItemRemovedEvent` events are fired respectively.

```
Entity Event Handler"
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

When the value of the ItemCount property changes, the `InventoryItemModifiedEvent` event occurs.

An event occurs if as follows.

- When the owner of the item is changed
- When the property value of the item type (in this example, TestItem) is changed

```
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

Your inventory remains after you leave the game. The list of items are checked in `InventoryItemInitEvent`, which occurs when entering the game.

```
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

A Component to support the key function for virtual operation controlling the player's movement in the mobile environment.

# Properties

| [AxisType](https://mod-developers.nexon.com/apiReference/Enums/AxisType) Axis |
| --- |
| Sets the axis type for the operation key to move. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) DownArrow |
| --- |
| Pushing the controller down triggers this key. |

| boolean DynamicStick |
| --- |
| The location of the control key moves to the user's touched location during play. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) LeftArrow |
| --- |
| Pushing the controller left triggers this key. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) RightArrow |
| --- |
| Pushing the controller right triggers this key. |

| [KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) UpArrow |
| --- |
| Pushing the controller up triggers this key. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of mapping the joystick’s downward position to the E key input. Movement is processed in the client space, so the key should also be changed in the client.

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

Vertical and horizontal movement and collision control for top-down gameplay with RectTiles. Gravity does not affect acceleration and deceleration. Used if the tile map is RectTile.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Acceleration ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is deprecated. Use the SpeedFactor property. |

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, the character on the rotating or slanted ladder will be affected by the shape of the ladder. When false, the character is not affected by the slant or rotation of the ladder. |

| boolean EnableJump ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shows whether you will use the jump feature. |

| boolean EnableShadow ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shows whether you will use the shadow. |

| boolean EnableTileCollision ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to use the collision feature with RectTileMap. There will be no collisions with collision tiles, and RectTileCollisionBeginEvent and RectTileCollisionEndEvent will not be caused if the value is false. |

| float JumpDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Jump speed reduction. The greater the value, the faster the fall. |

| float JumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The jump speed. The greater the value, the faster the jump. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the movement speed. The value multiplied by SpeedFactor will be the final speed.<br>If a player moves a directional key or calls MovementComponent:MoveToDirection() function, MoveVelocity's value will change. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ShadowColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The shadow's color. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ShadowOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The shadow's position. |

| float ShadowScalingRatio ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The rate of change of the shadow's size. The size varies depending on how high the Entity jumps. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ShadowSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The shadow's size. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) SpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When moving, this value is multiplied to the speed traveled in the X or Y direction. The greater the value, the faster the movement. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetGroundPosition() |
| --- |
| Returns position of the floor based on local coordinates. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetWorldGroundPosition() |
| --- |
| Returns position of the floor based on World coordinates. |

| boolean IsOnGround() |
| --- |
| Checks the current contact with the ground. Returns false when jumping. |

| void OnEnterRectTile([RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) enterEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is called when RectTileEnterEvent occurs. |

| void OnLeaveRectTile([RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) leaveEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is called when RectTileLeaveEvent is generated. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's position based on local coordinates. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets Entity's location to the world-based coordinates. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [KinematicbodyJumpEvent](https://mod-developers.nexon.com/apiReference/Events/KinematicbodyJumpEvent) |
| --- |
| This event occurs when changing the jumping state. |

| [RectTileCollisionBeginEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionBeginEvent) |
| --- |
| This event occurs when touching the collidable tile. |

| [RectTileCollisionEndEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionEndEvent) |
| --- |
| This event occurs when being freed from the collided tile. |

| [RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) |
| --- |
| This event occurs when entering a specific quadrangle tile. |

| [RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) |
| --- |
| This event occurs when leaving a specific quadrangle tile. |

# Examples

This is an example of leaping over a tile by using jump.

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
- [Control Character Movement from RectTileMap](/docs?postId=748)

Update 2025-08-27 PM 04:56


# LineGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Draws a line in UI and provides the function to set the properties of the line.

# Properties

| float Flexibility |
| --- |
| Adjusts the width of the bending part. It is applied when IsFlexible is true. |

| boolean IsFlexible |
| --- |
| Sets whether to draw the bending part continuously or separately. |

| boolean IsSmooth |
| --- |
| Draws a smooth curve if true, draws a sharp line if false. |

| boolean Loop |
| --- |
| If true, a closed line will be drawn automatically to connect the start and end points. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| [SyncList<LinePoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points |
| --- |
| The group of dots that form a line. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# LineRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Draws a line and provides the ability to set the properties of the line.

# Properties

| float Flexibility ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Adjusts the width of the bending part. It is applied when IsFlexible is true. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsFlexible ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to draw the bending part continuously or separately. |

| boolean IsSmooth ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Draws a smooth curve if true, draws a sharp line if false. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, a closed line will be drawn automatically connecting the start and end points. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material that the renderer will use. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same Layer. A higher number indicates higher priority. |

| [SyncList<LinePoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The group of dots that form a line. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

The following is an example of drawing a circle around an entity.

```
Function:   
[server only]
void OnBeginPlay ()
{
	self:DrawCircle(2, 20, Color.red, 1)
}
 
[server]
void DrawCircle ( number radius, integer vertexNum, Color color, number width )
{
	-- radius: radius
	-- vertexNum: number of vertices
	-- color: line color
	-- width: line thickness
 
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
- [Draw Various Lines](/docs?postId=749)

Update 2025-10-28 PM 02:21


# MapComponent

You can manage the map and change its unique properties.

# Properties

| float AirAccelerationXFactor |
| --- |
| Corrects the speed mid-air for Entities with RigidbodyComponents on the map. The higher the value is, the faster the speed in the air. |

| float AirDecelerationXFactor |
| --- |
| Adjusts how fast the X-axis movement stops if no input is received while Entities with RigidbodyComponents are in mid-air. |

| float FallSpeedMaxXFactor |
| --- |
| Corrects the X-axis maximum speed limit mid-air for Entities with RigidbodyComponents on the map. |

| float FallSpeedMaxYFactor |
| --- |
| Corrects the Y-axis maximum speed limit mid-air for Entities with RigidbodyComponents on the map. |

| float Gravity |
| --- |
| Corrects the gravity value of Entities with RigidbodyComponent on the map. |

| boolean IsDynamicMap ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether the map is dynamically created. |

| boolean IsInstanceMap |
| --- |
| Sets whether the instance map will be used. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LeftBottom ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The bottom-left value of the map area. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RightTop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The top-right of the map area. |

| [TileMapMode](https://mod-developers.nexon.com/apiReference/Enums/TileMapMode) TileMapMode ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Checks the tilemap mode of the map. |

| boolean UseCustomBound ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shows whether you will directly define and use the map area. You can define it using the LeftBottom and RightTop properties. When false, the area automatically created based on the map's form will be used as the map area. Works only when TileMapMode is MapleTile. |

| float WalkAccelerationFactor |
| --- |
| Sets the Factor for the movement speed of Entities with RigidbodyComponent within the map. The maximum speed cannot exceed the WalkSpeed for the RigidbodyComponent, which regulates the maximum speed. |

| float WalkDrag |
| --- |
| Corrects the tile friction of Entities with RigidbodyComponents on the map. The lower the value is, the better the tile slides. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| Vector2, Vector2 GetBound() |
| --- |
| Gets the map area composed of LeftBottom and RightTop. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# MapLayerComponent

Manages information for each layer of the map.

# Properties

| boolean IsVisible ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether the layer is rendered or not. Invisible in the Maker scene if false. |

| int32 LayerSortOrder ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The priority order of the layers. Smaller values are drawn in lower layers and behind them. |

| boolean Locked ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether the layer is locked. The maker prevents editing when locked. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MapLayerName ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The name of the map layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# MaskComponent

Makes only certain areas of the child UI Entity visible using a mask.

# Properties

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| Sets the free space on the top, bottom, left, and right of the mask. |

| [MaskShape](https://mod-developers.nexon.com/apiReference/Enums/MaskShape) Shape |
| --- |
| Specifies the shape of the mask. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) Softness |
| --- |
| Determines the area to blur naturally from the edge area of the mask. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

When initializing a UI Entity, padding is added to the mask, reducing the visible area. This is an example of setting `Softness` and processing the softness of the edge of the child entity.

```
Function:
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

Provides movement functions for controlling RigidbodyComponent, KinematicbodyComponent, and SideviewbodyComponent, with easy adjustments for jumping power and speed.

# Properties

| float InputSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the movement speed. The higher the value, the faster the speed.<br><br>- If RigidbodyComponent is true, affects horizontal speed.<br>- If KinematicMove is true, affects both vertical and horizontal speed.<br>- If KinamaticbodyComponent is true, affects both vertical and horizontal speed.<br>- If SideviewbodyComponent is true, affects both horizontal speed. |

| boolean IsClimbPaused ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Checks if paused while climbing. |

| float JumpForce ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the jump strength. The bigger the value is, the higher the jump. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| boolean DownJump() |
| --- |
| Jump down. Returns whether the jump down is successful or not. |

| boolean IsFaceLeft() |
| --- |
| Returns whether an Entity is pointing to the left. |

| boolean Jump() |
| --- |
| Jump. Returns whether the jump is successful or not. |

| void MoveToDirection([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) direction, float deltaTime) |
| --- |
| Moves towards the direction during the deltaTime. The unit of deltaTime is the second. Applied only when on the ladder. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's position based on local coordinates. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets Entity's location to the world-based coordinates. |

| void Stop() |
| --- |
| Stops moving. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ChangedMovementInputEvent](https://mod-developers.nexon.com/apiReference/Events/ChangedMovementInputEvent) |
| --- |
| Occurs when the movement input changes on the MovementComponent. |

| [ClimbPauseEvent](https://mod-developers.nexon.com/apiReference/Events/ClimbPauseEvent) |
| --- |
| This event is fired when riding on an object and stopping. |

# Examples

When the player looks left, movement starts automatically. This example also shows how to adjust jump height, initiate jumping, and stop when touching a specific entity.

```
[Sync]
boolean isStarted = false
[Sync]
boolean isFinished = false 

Function:
[client only]
void OnUpdate ( number delta)
{
	if self.isFinished then
		self.Entity.MovementComponent:Stop() -- Movement by input is also impossible
		return
	end
	 
	if self.isStarted == false and self.Entity.MovementComponent:IsFaceLeft() then
		self.isStarted = true;
	end
 
	if self.isStarted == false then
		return
	end
 
self.Entity.MovementComponent:MoveToDirection(Vector2(1,0), delta)
}

Entity Event Hanlder:
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
		self.isFinished = true
	end
}
```

# SeeAlso

- [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent)
- [Entity Movement Control Using MovementComponent](/docs?postId=546)

Update 2025-10-28 PM 02:21


# NameTagComponent

Displays the Entity's name tag and set-related information.

# Properties

| boolean Bold ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to use bold text. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Changes the text color. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) FontOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Adjusts position of the text within the name tag. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the font size. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Name ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the name on the name tag. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) NameTagRUID |
| --- |
| You can change the name tag shape. |

| float OffsetY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets Offset of the name tag's position. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of changing the NameTag to prevent users from knowing each other's nicknames. UserId is used to determine the word and color of the NamgTag when a user enters.

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
- [Naming Entities](/docs?postId=29)

Update 2025-08-27 PM 04:56


# PhysicsColliderComponent

Sets the size and offset of the physics rigid body and raises an event when it collides with another Entity. This is automatically connected to the PhysicsRigidbodyComponent of the parent Entity or its own Entity. If PhysicsRigidbodyComponent does not exist, it is treated as a static physics rigid body.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Box Collider size. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the radius length of the Circle Collider. |

| boolean ClientOnly ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| If true, physics operations occur in the Client, and functions and Property changes are possible in the Client space. Physics calculation results are not synchronized with other Clients. If false, function use and Property change are possible in the Server space. Physics calculation results are synchronized with other Clients. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Collider's offset. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Collider's shape. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets up a collision group that can determine if there is a collision or not. |

| float Density ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the density value to use when UseDensity of the attached PhysicsRigidbodyComponent is true. |

| boolean EnableContactEvent ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether PhysicsContactBeginEvent and PhysicsContactEndEvent occur. If false, neither event occurs. |

| float Friction ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the coefficient of friction. Must set a value equal to or greater than 0. Valid when UseCustomPhysicalProperties is true. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether to support the legacy system. ColliderOffset affects the Entity's SpriteRendererComponent in the previous system. The previous system is no longer supported and will be deleted at a later date. |

| boolean IsSensor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, no physical interaction will occur, but a collision event will occur. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The positions of the points that make up the polygonal colliding body. Valid when ColliderType is Polygon. |

| float Restitution ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines how much bounce occurs after a collision. Must set a value equal to or greater than 0. Valid when UseCustomPhysicalProperties is true. |

| boolean UseCustomPhysicalProperties ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Friction and Restituion values are applied to physics rigid bodies. If false, the value of the connected PhysicsRigidbodyComponent is used. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [PhysicsRigidbodyComponent](https://mod-developers.nexon.com/apiReference/Components/PhysicsRigidbodyComponent) GetAttachedPhysicsRigidbody() |
| --- |
| Returns the connected PhysicsRigidbodyComponent. |

| void OnContactBegin([PhysicsContactBeginEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactBeginEvent) beginEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Called when two physics rigid bodies begin to collide. |

| void OnContactEnd([PhysicsContactEndEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactEndEvent) endEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Called when contact ends. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [PhysicsContactBeginEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactBeginEvent) |
| --- |
| This event is raised when two physics rigid bodies begin to collide. |

| [PhysicsContactEndEvent](https://mod-developers.nexon.com/apiReference/Events/PhysicsContactEndEvent) |
| --- |
| This event is raised when two physics rigid bodies end collision. |

# Examples

This example shows pressing P on the keyboard changes the CollisionGroup to Monster, preventing it from colliding with other Monsters. The SleepingMode must be set to NeverSleep.

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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-12-02 PM 01:55


# PhysicsRigidbodyComponent

Entity is controlled by the physics engine. You can set values that affect physics calculations.

# Properties

| float AngularDamping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the coefficient of resistance to changes in angular velocity. The larger the value is, the more force is required to change the angular velocity. Must set a value equal to or greater than 0. |

| [BodyType](https://mod-developers.nexon.com/apiReference/Enums/BodyType) BodyType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the rigid body type. |

| [PhysicsCollisionDetectionMode](https://mod-developers.nexon.com/apiReference/Enums/PhysicsCollisionDetectionMode) CollisionDetectionMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the detection method for physics collision. |

| boolean FixedRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the object does not rotate due to physics interactions. |

| float Friction ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the coefficient of friction. The smaller the value is, the better it slides. Must set a value equal to or greater than 0. |

| float GravityScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets how much the effect of gravity is. If the value is 0, gravity has no effect and if it is 1, gravity has a direct effect. The larger the value is, the more effect gravity has. |

| float LinearDamping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the coefficient of resistance for the change of the linear velocity. The larger the value is, the more force is required to change the linear velocity. |

| float Mass ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the mass value. Calculate the density value using the values of the area and mass of the Collider. |

| float Restitution ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines how much bounce occurs after a collision. 0 is close to perfectly inelastic collision, 1 is close to perfectly elastic collision. |

| [PhysicsSleepingMode](https://mod-developers.nexon.com/apiReference/Enums/PhysicsSleepingMode) SleepingMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the sleep state of the collision. |

| boolean UseDensity ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the mass value is calculated using the density and collider values of PhysicsRigidbodyComponent instead of the mass value.<br>If multiple PhysicsRigidbodyComponents are attached to one PhysicsRigidbodyComponent, the value obtained by multiplying the density of each PhysicsRigidbodyComponent by the collider size is used as the mass value. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ApplyAngularImpulse(float impulse) |
| --- |
| Applies each impulse to the rigid body counterclockwise (CCW). |

| void ApplyForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) force) |
| --- |
| Apply a force in a specific direction to the center of gravity of a rigid body. |

| void ApplyForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) force, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| Apply a force to a rigid body in a specific direction at a specific location in the world. |

| void ApplyLinearImpulse([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) impulse) |
| --- |
| Applies a linear impulse in a specific direction to the center of gravity of a rigid body. |

| void ApplyLinearImpulse([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) impulse, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| Applies a linear impulse to a rigid body in a specific direction at a specific location in the world. |

| void ApplyTorque(float force) |
| --- |
| Applies a rotational force to the rigid body in a counterclockwise (CCW) direction. |

| void ClearPhysicsOwnership() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| Automatically grants physics ownership. A Contact Event occurs only on the Server. |

| float GetAngularVelocity() |
| --- |
| Returns the angular velocity of a rigid body. |

| float GetDensity() |
| --- |
| Returns the density of a rigid body. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetLinearVelocity() |
| --- |
| Returns the linear velocity of a rigid body. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetLinearVelocityAtPoint([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) worldPoint) |
| --- |
| Returns the linear velocity of a rigid body at a specific location in the world. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetPosition() |
| --- |
| Returns the location of a rigid body. |

| float GetRotation() |
| --- |
| Returns the angle of rotation of a rigid body. |

| void SetAngularVelocity(float velocity) |
| --- |
| Sets the angular velocity of a rigid body. |

| void SetClientAsPhysicsOwner([string](https://mod-developers.nexon.com/apiReference/Lua/string) userId) ![custom](https://img.shields.io/static/v1?label=&amp;message=ServerOnly&amp;color=mediumvioletred) |
| --- |
| Grants physics ownership to a specific user. A Contact Event occurs on the corresponding Client and the Server. |

| void SetLinearVelocity([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) velocity) |
| --- |
| Sets the linear velocity of a rigid body. Not affected by mass. |

| void SetLinearVelocityX(float velocityX) |
| --- |
| Sets the linear velocity X value of the rigid body. It is not affected by mass. |

| void SetLinearVelocityY(float velocityY) |
| --- |
| Sets the linear velocity Y value of the rigid body. It is not affected by mass. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets the location of the rigid body. |

| void SetRotation(float angle) |
| --- |
| Returns the angle of rotation of a rigid body. |

| void SetServerAsPhysicsOwner() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| Grants physics ownership to the server. A Contact Event occurs on the Server. |

| void Sleep() |
| --- |
| Puts the rigid body to sleep. |

| void Wake() |
| --- |
| Change a rigid body to an awake state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

#### ApplyLinearImpulse

```
Event Handler:
[self] HandlePhysicsContactBeginEvent ( PhysicsContactBeginEvent event )
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
HandlePhysicsContactBeginEvent ( PhysicsContactBeginEvent event )
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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-12-02 PM 01:55


# PhysicsSimulatorComponent

This handles the map’s physics simulation and sets values for overall physics operations.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Gravity ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the gravity value for physics calculations. |

| boolean Paused ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to pause the physics calculations. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) WorldBounds ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets the range for physics calculations. If Entity is out of range, physics calculations will no longer occur. For example, if the value is set to (100, 100), a 200 x 200 rectangular area will be the physics calculations range from center of axis (0,0). |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void SetPositionIteration(int32 count) |
| --- |
| Sets the number of repetitions of speed calculations by the physics engine. As the value increases, the physics calculation result will be more elaborate, and more calculation time will be taken. The minimum value is 1 and the maximum value is 300. It does not apply to server operations. |

| void SetVelocityIteration(int32 count) |
| --- |
| Sets the number of repetitions of the location calculations by the physics engine. As the value increases, the physics calculation result will be more elaborate, and more calculation time will be taken. The minimum value is 1 and the maximum value is 30. It does not apply to server operations. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This example shows pausing the physics simulation or changing the gravity settings based on the key pressed.

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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-10-28 PM 02:21


# PixelGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Designates pixel values to enable the player to draw sprites on the UI as desired. We recommend using a size of less than 16x16.

# Properties

| int32 Height ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The height of the sprite. |

| boolean IgnoreMapLayerCheck |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| int32 OrderInLayer |
| --- |
| Determines the priority within the same Layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| boolean RaycastTarget |
| --- |
| This can be tapped or clicked when it's set as true. The UI hidden behind it will not respond to the input. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| When two or more entities overlap, the priority will be determined by the Sorting Layer. |

| int32 Width ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The width of the sprite. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void FillColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Changes the color of the entire sprite. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) GetPixel(int32 x, int32 y) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the pixel value for the desired location. Bottom left is (1, 1). |

| [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixels() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns all pixel values for the current sprite. Row comes first. (Row-major) |

| [table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixelsAsRGBAInt() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the entire pixel value as an RGBA integer. Row-major order. |

| void ResetWithColor(int32 width, int32 height, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Resets the sprite to be customized to the size you've entered and fills it with the color. |

| void ResetWithColors(int32 width, int32 height, [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Resets the sprite to be customized to the size you've entered and sets the values of the pixels according to the values of the tables you've entered. The number of table elements must be the same as the number of all pixels (Width*Height). Row comes first. (Row-major) |

| void SetAlpha(float alpha) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Changes the alpha value for the entire sprite. |

| void SetPixel(int32 x, int32 y, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the pixel value for the location you've entered. Bottom left is (1, 1). |

| void SetPixels([table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the values of the pixels according to the values of tables you've entered. The number of table elements must be the same as the number of all pixels (Width*Height). Row comes first. (Row-major) |

| void SetPixelsByRGBAInt([table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the entire pixel value as an RGBA integer. The table length must match the total pixel count (Width*Height) and is row-major. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of how to make pixels blink randomly when you press the F key on the keyboard.

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

Designates colors by each pixel to provide a feature of creating sprites as desired. We recommend using a size of less than 16x16.

# Properties

| int32 Height ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The height of the sprite. |

| boolean IgnoreMapLayerCheck |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| int32 OrderInLayer |
| --- |
| Determines the priority within the same Layer. A higher number indicates higher priority. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| int32 Width ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The width of the sprite. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void FillColor([Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Fills the whole sprite with the color you've entered. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) GetPixel(int32 x, int32 y) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the pixel value of the location you've entered. Bottom left is (1, 1). |

| [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixels() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the values of all pixels. Row comes first. (Row-major) |

| [table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) GetPixelsAsRGBAInt() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the entire pixel value as an RGBA integer. The table length must match the total pixel count (Width*Height) and is row-major. |

| void ResetWithColor(int32 width, int32 height, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Resets the sprite to be customized to the size you've entered and fills it with the color. |

| void ResetWithColors(int32 width, int32 height, [table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Resets the sprite to be customized to the size you've entered and sets up values of the pixels with the values of tables you've entered. The number of table elements must be the same as the number of all pixels (Width * Height). Row comes first. (Row-major) |

| void SetAlpha(float alpha) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the alpha value of the sprite. |

| void SetPixel(int32 x, int32 y, [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) color) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the pixel value of the location you've entered. Bottom left is (1, 1). |

| void SetPixels([table<Color>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Sets the values of pixels with the values of tables you've entered. The number of table elements must be the same as the number of all pixels (Width * Height). Row comes first. (Row-major) |

| void SetPixelsByRGBAInt([table<integer>](https://mod-developers.nexon.com/apiReference/Lua/table) pixels) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the entire pixel value as an RGBA integer. Row-major order. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of randomly blinking pixels when the F key on the keyboard is pressed.

```
Property:
[None]
table<Vector2> puzzles
[None]
PixelRendererComponent pixelRendererComp = nil
[None]
table<Vector> puzzlePos

Method:
[client only]
void OnBeginPlay ()
{
	self.pixelRendererComp = self.Entity.PixelRendererComponent
	self.pixelRendererComp:FillColor(Color.black)
	 
	for i=1, 9 do
		local xPos = _UtilLogic:RandomIntegerRange(1, 3)
		local yPos = _UtilLogic:RandomIntegerRange(1, 3)
		self.puzzles[i] = Vector2(xPos, yPos)
	end
}

[client only]
void StartPuzzle ( number index )
{
	if index > #self.puzzles then return end
	 
	self.pixelRendererComp:SetPixel(self.puzzles[index].x, self.puzzles[index].y, Color.yellow)
	wait(0.35)
	self.pixelRendererComp:SetPixel(self.puzzles[index].x, self.puzzles[index].y, Color.black)
	wait(0.1)
	self:StartPuzzle(index + 1)
}

Event Handler:
[service] [InputService]
HandleKeyDownEvent ( KeyDownEvent event )
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
- [Setting Sprite Color by Pixel](/docs?postId=693)

Update 2025-12-02 PM 01:55


# PlayerComponent

Represents a player and provides related function.

# Properties

| integer Hp ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Current Hp. |

| integer MaxHp ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Maximum Hp. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Nickname ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Nickname of the player. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ProfileCode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The player's profile code. |

| boolean PVPMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether players can attack each other. |

| float RespawnDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The amount of time to respawn after dying. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) RespawnPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the location to respawn. If not specified, the SpawnLocation is set as the top priority while the point of the map entrance is set as the second priority. |

| number RespawnTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The expected time to respawn. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) UserId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| A unique identifier for the player. It can be used in the targetUserId parameter of the Client execution space control function. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| boolean IsDead() |
| --- |
| Returns whether the player is dead or not. |

| void MoveToEntity([string](https://mod-developers.nexon.com/apiReference/Lua/string) entityID) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Moves the player to the same location as an Entity corresponding to the entityID or to another map if the entityID belongs to it. |

| void MoveToEntityByPath([string](https://mod-developers.nexon.com/apiReference/Lua/string) worldPath) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Moves the player to the same location as an Entity in the worldPath, or another map if the worldPath points to it. |

| void MoveToMapPosition([string](https://mod-developers.nexon.com/apiReference/Lua/string) mapID, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) targetPosition) ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Moves the player to a specific location on a specific map. |

| void ProcessDead([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Causes the player to die. |

| void ProcessRevive([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Revives player. |

| void Respawn() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Performs a respawn. For the Respawn position, RespawnPosition is the first priority, SpawnLocation in the map is the second, and the point of map entrance is the third. |

| void SetPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) position) |
| --- |
| Sets an Entity's position based on local coordinates. |

| void SetWorldPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPosition) |
| --- |
| Sets the Entity position in world coordinates. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

Changes the respawn position to the object's position while touching an object named "Check Point". To see if the respawn position has changed, touch an object named "Do Not Touch", which kills the player.

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
- [Setting and Controlling Player](/docs?postId=547)

Update 2025-10-28 PM 02:21


# PlayerControllerComponent

A component related to the Player control. Associates inputs with actions and controls their flow.

# Properties

| boolean AlwaysMovingState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether to always play the walk animation. |

| int32 FixedLookAt ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Fixes direction of view when moving. |

| float LookDirectionX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The direction the character is facing currently based on the axis X. Right if positive, left if negative. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ActionAttack() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Attack Key. |

| void ActionCrouch() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Crouch Key. |

| void ActionDownJump() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action on ActionDownJump. |

| void ActionEnterPortal() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Portal Key. |

| void ActionInteraction([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key, boolean isKeyDown) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Interaction Key. |

| void ActionJump() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Jump Key. |

| void ActionSit() ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| The action upon entering the Sit Key. |

| void AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName, func -> boolean conditionFunction) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Add conditions to trigger an action. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetActionName([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the names of the actions mapped to the keys. |

| void RemoveActionKey([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Removes the actions connected to the defined keys. |

| void RemoveAllActionKeyByActionName([string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Removes all actions connected to the defined names. |

| void SetActionKey([KeyboardKey](https://mod-developers.nexon.com/apiReference/Enums/KeyboardKey) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) actionName, func -> boolean conditionFunction = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Map actions to the keys. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ChangedLookAtEvent](https://mod-developers.nexon.com/apiReference/Events/ChangedLookAtEvent) |
| --- |
| This occurs when changing the direction for the character to face in the PlayerControllerComponent. |

| [PlayerActionEvent](https://mod-developers.nexon.com/apiReference/Events/PlayerActionEvent) |
| --- |
| This event is raised when the player uses an Action. |

# Examples

In this example, keyboard inputs are linked to predefined actions, such as Jump, Portal, Crouch, Attack, and Sit. User-defined action names are linked to keyboard inputs for logging.

Tap the B key to Attack and the N key to Jump. You can also tap the G key to view user-defined action names being output to the log.

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
- [Setting and Controlling Player](/docs?postId=547)

Update 2025-08-27 PM 04:56


# PolygonGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

This feature allows you to draw polygons in UI.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| The color of the polygon. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points |
| --- |
| The group of dots that form a polygon. |

| boolean UseCustomUVs |
| --- |
| If true, the UV value ​​set in the UVs is applied to the shape. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) UVs |
| --- |
| Sets the vertex's UV. It must have the same length as Points. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| boolean IsDrawable() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Connect dots saved in Points property to see if they make up a polygon. The polygon will not be completed if any lines cross with one another. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# PolygonRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

This feature allows you to draw polygons.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color of the polygon. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material that will be used at the renderer. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Points ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The group of dots that form a polygon. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| boolean UseCustomUVs ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the UV value ​​set in the UVs is applied to the shape. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) UVs ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the vertex's UV. It must have the same length as Points. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replace the material to be applied to the renderer. |

| boolean IsDrawable() |
| --- |
| Connect dots saved in Points property to see if they make up a polygon. The polygon will not be completed if any lines cross with one another. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# SeeAlso

- [Draw Polygon](/docs?postId=1080)

Update 2025-08-27 PM 04:56


# PortalComponent

Sets the movement between Portals. You can move from one Portal to another.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the center point position of the collider rectangle based on Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the collider rectangle. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Collision group in Portal. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether this Component will operate with the previous system. In the new system, the Collider is affected by TransformComponent's rotation and size. |

| [EntityRef](https://mod-developers.nexon.com/apiReference/Misc/EntityRef) PortalEntityRef |
| --- |
| Sets a destination portal. Only an entity that has PortalComponent can be set as a destination. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [PortalUseEvent](https://mod-developers.nexon.com/apiReference/Events/PortalUseEvent) |
| --- |
| This event occurs when a player uses a portal. |

# Examples

Place at least 3 portals on the map. Adding component to portal will randomly change the connected portal every second.

```
Method:
[client only]
void OnBeginPlay ()
{
	-- Traverses all Entities in the current map and saves EntityRef of Entities with PortalComponent
	self._T.portalRefList = {}
	local function findAndAddPortal(parent)
		-- Except itself
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
	      
	-- Returns if there is no other portal except itself
	if #self._T.portalRefList < 1 then
		return
	end
	     
	-- Sets the timer that randomly changes the connected portal every second.
	_TimerService:SetTimerRepeat(function()
		self.Entity.PortalComponent.PortalEntityRef = self._T.portalRefList[_UtilLogic:RandomIntegerRange(1, #self._T.portalRefList)]
	end, 1)
}
```

# SeeAlso

- [math](https://mod-developers.nexon.com/apiReference/Lua/math)
- [table](https://mod-developers.nexon.com/apiReference/Lua/table)
- [TimerService](https://mod-developers.nexon.com/apiReference/Services/TimerService)
- [Making a Portal to Move to Another Location](/docs?postId=90)

Update 2025-08-27 PM 04:56


# PrismaticJointComponent

You can create and delete PrismaticJoint. Constrains a connected rigid body to move only in specific axis direction.

# Properties

| [SyncList<PrismaticJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Set Joint information. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAxis) |
| --- |
| Add Joint. Return index on success, -1 on failure |

| void DestroyJoint(int32 index) |
| --- |
| Removes the Joint whose sequence number corresponds to the index. |

| int32 GetJointsCount() |
| --- |
| Return the number of joints. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| Set the CollideConnected value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| Set the LocalAnchorA value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Set the LocalAnchorA value of the Joint whose sequence number corresponds to index. |

| void SetLocalAxis(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAxis) |
| --- |
| Set the LocalAxis value of the Joint whose sequence number corresponds to index. |

| void SetLowerTranslation(int32 index, float lowerTranslation) |
| --- |
| Set the LowerTranslation value of the Joint whose sequence number corresponds to index. |

| void SetMaxMotorForce(int32 index, float maxMotorForce) |
| --- |
| Sets the MaxMotorForce value of the Joint whose sequence number corresponds to index. |

| void SetMotorEnable(int32 index, boolean enable) |
| --- |
| Set the MotorEnable value of the Joint whose sequence number corresponds to index. |

| void SetMotorSpeed(int32 index, float speed) |
| --- |
| Set the MotorSpeed value of the Joint whose sequence number corresponds to index. |

| void SetUpperTranslation(int32 index, float upperTranslation) |
| --- |
| Set the UpperTranslation value of the Joint whose sequence number corresponds to index. |

| void SetUseLimits(int32 index, boolean useLimits) |
| --- |
| Set the UseLimits value of the Joint whose sequence number corresponds to index. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

In this example, pressing the W and S keys on the keyboard changes the SetMotorSpeed value. The Entity set as TargetEntityRef increases and decreases equal to the value of UpperTranslation and LowerTranslation of the PrismaticJointComponent.

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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-10-28 PM 02:21


# PulleyJointComponent

Create and delete PulleyJoint. Two connected rigid bodies move like a pulley.

# Properties

| [SyncList<PulleyJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Set Joint information. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) groundAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) groundAnchorB, float ratio = 1) |
| --- |
| Add Joint. Return index on success, -1 on failure. |

| void DestroyJoint(int32 index) |
| --- |
| Removes the Joint whose sequence number corresponds to the index. |

| int32 GetJointsCount() |
| --- |
| Returns the number of joints. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| Set the CollideConnected value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| Set the LocalAnchorB value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Set the LocalAnchorB value of the Joint whose sequence number corresponds to index. |

| void SetRatio(int32 index, float ratio) |
| --- |
| Set the Ratio value of the Joint whose sequence number corresponds to index. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# SeeAlso

- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-10-28 PM 02:21


# RawImageGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the ability to output RawImage to the UI.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| void SetRawImage([RawImage](https://mod-developers.nexon.com/apiReference/Misc/RawImage) rawImage) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Draws the RawImage. Does not draw anything if nil is passed. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# RawImageRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides a function to output RawImage.

# Properties

| boolean IgnoreMapLayerCheck |
| --- |
| Does not automatically replace when designating the Map Layer name into the SortingLayer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| int32 OrderInLayer |
| --- |
| Determines the priority within the same Layer. A higher number indicates higher priority. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer |
| --- |
| When two or more Entities overlap, the priority is determined by the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

| void SetRawImage([RawImage](https://mod-developers.nexon.com/apiReference/Misc/RawImage) rawImage) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Draws the RawImage. Does not draw anything if nil is passed. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

Update 2025-08-27 PM 04:56


# RectTileMapComponent

Provides a rectangular tile map function.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GridSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The size of the grid. |

| boolean IgnoreMapLayerCheck |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsOddGridPosition |
| --- |
| Positions the tilemap off of the grid's base point. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same Layer. A greater number indicates higher priority. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| If true, it may collide with a Dynamic rigid body (PhysicRigidbody) using the Physics feature. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| int32 TileCount ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Total number of tiles currently placed. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) TileSetRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The RUID of the tileset. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void BoxFill(int32 tileIndex, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| Places tiles in a quadrangle area "from" to "to". tileIndex is the index number of the tile. |

| void BoxFill([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| Places tiles in a quadrangle area from "from" to "to". tileName is the Name property of the tile. |

| void BoxRemove([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) from, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) to) |
| --- |
| Removes tiles in the quadrangle area "from" to "to". |

| void Clear() |
| --- |
| Removes all tiles. |

| [List<Vector2Int>](https://mod-developers.nexon.com/apiReference/Misc/List-1) GetAllTilePositions() |
| --- |
| Returns a list of all tile locations in the current tilemap. |

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetTile(int32 cellPositionX, int32 cellPositionY) |
| --- |
| Returns the tile information at the corresponding location. |

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetTile([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| Returns the tile information at the corresponding location. |

| void RemoveTile(int32 cellPositionX, int32 cellPositionY) |
| --- |
| Removes the tile at the corresponding location. |

| void RemoveTile([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| Removes the tile at the corresponding location. |

| void Reset() |
| --- |
| Resets tile map to the initial state. |

| void SetTile(int32 tileIndex, int32 cellPositionX, int32 cellPositionY) |
| --- |
| Places the tile in the corresponding location with tileIndex being the index number of the tile. |

| void SetTile(int32 tileIndex, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| Places the tile in the corresponding location with tileIndex being the index number of the tile. |

| void SetTile([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, int32 cellPositionX, int32 cellPositionY) |
| --- |
| Places the tile in the corresponding location with tileName being the Name property of the tile. |

| void SetTile([string](https://mod-developers.nexon.com/apiReference/Lua/string) tileName, [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| Places the tile in the corresponding location with tileName being the Name property of the tile. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) ToCellPosition(float worldPositionX, float worldPositionY) |
| --- |
| Converts the coordinates of real number world space to those of integer tile map space. |

| [Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) ToCellPosition([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPosition) |
| --- |
| Converts the coordinates of real number world space to those of integer tile map space. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPosition(int32 cellPositionX, int32 cellPositionY) |
| --- |
| Converts the coordinates of integer tile map space to those of real number world space. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPosition([Vector2Int](https://mod-developers.nexon.com/apiReference/Misc/Vector2Int) cellPosition) |
| --- |
| Converts the coordinates of integer tile map space to those of real number world space. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

# Examples

This is an example where player movement speed changes depending on the tile stepped on.

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
	 
	-- Convert the world space coordinates to TileMap space coordinates
	local cellPos = tilemap:ToCellPosition(worldPos)
	  
	-- Inquire about tile information in the current location
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
- [Creating Maps with SideViewRectTile Mode](/docs?postId=758)
- [Utilizing RectTileMap](/docs?postId=589)

Update 2025-10-28 PM 02:21


# RevoluteJointComponent

Create and delete RevoluteJoint. Controls the relative rotation of connected rigid bodies.

# Properties

| [SyncList<RevoluteJoint>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Joints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Set Joint information. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| int32 AddJoint([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) targetEntity, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Add Joint. Return index on success, -1 on failure |

| void DestroyJoint(int32 index) |
| --- |
| Removes the Joint whose sequence number corresponds to the index. |

| int32 GetJointsCount() |
| --- |
| Returns the number of joints. |

| void SetCollideConnected(int32 index, boolean collideConnected) |
| --- |
| Set the CollideConnected value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorA(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorA) |
| --- |
| Set the LocalAnchorA value of the Joint whose sequence number corresponds to index. |

| void SetLocalAnchorB(int32 index, [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) localAnchorB) |
| --- |
| Set the LocalAnchorB value of the Joint whose sequence number corresponds to index. |

| void SetLowerAngle(int32 index, float lowerAngle) |
| --- |
| Set the LowerAngle value of the Joint whose sequence number corresponds to index. |

| void SetMaxMotorTorque(int32 index, float maxMotorTorque) |
| --- |
| Sets the MaxMotorTorque value of the Joint whose sequence number corresponds to index. |

| void SetMotorEnable(int32 index, boolean enable) |
| --- |
| Set the MotorEnable value of the Joint whose sequence number corresponds to index. |

| void SetMotorSpeed(int32 index, float speed) |
| --- |
| Set the MotorSpeed value of the Joint whose sequence number corresponds to index. |

| void SetUpperAngle(int32 index, float upperAngle) |
| --- |
| Set the UpperAngle value of the Joint whose sequence number corresponds to index. |

| void SetUseLimits(int32 index, boolean useLimits) |
| --- |
| Set the UseLimits value of the Joint whose sequence number corresponds to index. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

#### SetMotorSpeed

In this example, tap the W or S keys on the keyboard to rotate the Entity set as TargetEntityRef.

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
- [Applying Physics to Entities](/docs?postId=761)
- [Using Physics](/docs?postId=757)
- [Using Various Physics Joints](/docs?postId=760)

Update 2025-10-28 PM 02:21


# RigidbodyComponent

Applies MapleStory movements. These will be influenced by gravity and acceleration and deceleration effects.

# Properties

| float AirAccelerationX |
| --- |
| Calibrates speed in midair. The higher the value, the faster the speed. |

| float AirDecelerationX |
| --- |
| Calibrates how quickly x-axis movement speed stops when there is no input while in midair. |

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, the character on the rotating or slanted ladder will be affected by the shape of the ladder. When false, the character is not affected by the slant or rotation of the ladder. |

| float DownJumpSpeed |
| --- |
| Controls the speed of bouncing up when jumping down. |

| boolean EnableKinematicMoveJump |
| --- |
| If KinematicMove is true, jumping can be allowed or not allowed. |

| float FallSpeedMaxX |
| --- |
| Calibrates max x-axis speed limit while in the air. |

| float FallSpeedMaxY |
| --- |
| Calibrates max y-axis speed limit while in the air. |

| float Gravity |
| --- |
| Gravity value. Related to the drop velocity when moving mid-air. The bigger the input value, the faster the drop velocity is. |

| boolean IgnoreMoveBoundary |
| --- |
| It is possible to leave the terrain-created map area if true. |

| boolean IsBlockVerticalLine |
| --- |
| Vertical topography is unconditionally blocked if true, unlike basic movement. This is used to make it impossible to pass through topography like walls. |

| boolean IsolatedMove |
| --- |
| Does not fall even when reaching the end of the foothold if the value is true. You can leave the foothold with external movements, such as jumps. |

| float JumpBias |
| --- |
| Sets the character's hang time when jumping. |

| boolean KinematicMove |
| --- |
| If true, movement will change to top-down vertical and horizontal movement. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) KinematicMoveAcceleration |
| --- |
| Set the movement speed. Operates if KinematicMove is true. |

| [AutomaticLayerOption](https://mod-developers.nexon.com/apiReference/Enums/AutomaticLayerOption) LayerSettingType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the relationship between the RigidbodyComponent and SortingLayer values of foothold, ladder, and rope. |

| float Mass |
| --- |
| Sets the mass of the object. The larger the value is, the slower the acceleration and deceleration, and the lower the reactivity to external factors. Must set a value equal to or greater than 0. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The input required for the move. Inputs are mainly controlled by MovementComponent. A positive X indicates right, and a positive Y indicates upward. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RealMoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates the amount shifted just before. It is read-only and is used when movement amount information is required. A valid execution space exists according to the moving subject, and LocalPlayer has the values from the Client and others from the Server. |

| float WalkAcceleration |
| --- |
| It indicates the acceleration/deceleration value when moving the terrain. The bigger the input, the faster it is to reach the maximum speed. |

| float WalkDrag |
| --- |
| The force to resist slipping when moving on the terrain. The larger the input, the faster it stops without slipping. The final applied values range from 0.5 to 2 by calculating the map, terrain, and character attributes. |

| float WalkJump |
| --- |
| A value for the height of jump. The bigger the value is, the higher the jump. |

| float WalkSlant |
| --- |
| A value associated with the capability to climb slopes when moving on the terrain. Can go over steep slopes if the input is large enough. The value is between 0 and 1. |

| float WalkSpeed |
| --- |
| Adjusts the maximum movement speed value when moving on the terrain. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void AddForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| Adds the force applied to Entity. The Entity will have accelerating and decelerating movements in the direction of the force that was added to the existing force. |

| void AttachTo([string](https://mod-developers.nexon.com/apiReference/Lua/string) entityId, [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) offset) |
| --- |
| Attaches this Entity to the Entity corresponding to entityId. With this, this Entity doesn't have a physical motion and becomes subordinate to the movement of the attached Entity. |

| void Detach() |
| --- |
| Takes off the Entity which was attached to another Entity using RigidbodyComponent:AttachTo(string, Vector3). |

| boolean DownJump() |
| --- |
| Performs jumping down. Jumping down is only valid on terrain. |

| [Foothold](https://mod-developers.nexon.com/apiReference/Misc/Foothold) GetCurrentFoothold() |
| --- |
| Returns the Foothold that you are currently stepping on. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetCurrentFootholdPerpendicular() |
| --- |
| Returns the vertical line of the terrain being treaded on. |

| boolean IsOnGround() |
| --- |
| Checks whether standing on the terrain or not. |

| boolean JustJump([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) jumpRate) |
| --- |
| Makes a target jump. |

| void PositionReset() |
| --- |
| The accumulated information of calculating location is deleted and a new calculation is made based on the current location. |

| boolean PredictFootholdEnd(float distance, boolean isFoward) |
| --- |
| Checks whether you can move the distance from the step you are currently stepping on. If isForward is true, checks the right direction, if it is false, checks the left direction.<br>Returns true if the distance from the current position to the end of the foothold is farther than distance, and false if it is closer. Returns false if you are not stepping on a foothold. |

| void SetForce([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| Sets the force applied to the Entity. The Entity will have accelerating and decelerating movements in the direction of the force set. |

| void SetForceReserve([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) forcePower) |
| --- |
| Instead of applying the force immediately, it replaces the force with the given input after finishing the move in the current frame. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's position based on local coordinates. |

| void SetUseCustomMove(boolean isUse) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is deprecated. Use the Enable property of RigidbodyComponent. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's location to the world-based coordinates. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [FootholdCollisionEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdCollisionEvent) |
| --- |
| This event occurs when RigidBodyComponent collides with the foothold. |

| [FootholdEnterEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdEnterEvent) |
| --- |
| This event occurs when the entity is attached to the Foothold. |

| [FootholdLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/FootholdLeaveEvent) |
| --- |
| This event occurs when the entity comes off of the Foothold. |

| [RigidbodyAttachEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyAttachEvent) |
| --- |
| The event occurs when the Entity attaches to another Entity using RigidbodyComponent:AttachTo(string, Vector3). Players occur on the client, other entities occur on the server. |

| [RigidbodyClimbableAttachStartEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyClimbableAttachStartEvent) |
| --- |
| This Event occurs before the Avatar climbs the ladder or rope. |

| [RigidbodyClimbableDetachEndEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyClimbableDetachEndEvent) |
| --- |
| Event that occurs after an Avatar falls from a ladder or rope. |

| [RigidbodyDetachEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyDetachEvent) |
| --- |
| This event is raised when the entity is released from its Attach state via RigidbodyComponent:Detach(). Players occur on the client and other entities occur in server space. |

| [RigidbodyKinematicMoveJumpEvent](https://mod-developers.nexon.com/apiReference/Events/RigidbodyKinematicMoveJumpEvent) |
| --- |
| This event occurs when jumping or landing while the KinematicMove property is true. |

# Examples

#### AttachTo

Attaches when touching a specific object, and detaches after 3 seconds of being attached.

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

Determines whether the end point of the foothold is separated by a specific distance. Returns true if the end of the foothold is closer than 10 in the positive direction from the foothold the character is stepping on, and false otherwise. The Enable state of the Entity is then changed depending on the results.

```
Function:
[client only]
void OnUpdate(number delta)
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
- [Understanding MapleStory Movement Concepts](/docs?postId=750)

Update 2025-12-02 PM 01:55


# ScrollLayoutGroupComponent

You can control the components related to the scroll view as one.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) CellSize |
| --- |
| Fixed size of child UI Entity. Grid type only. |

| [ChildAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/ChildAlignmentType) ChildAlignment |
| --- |
| Set the alignment direction of the child UI Entity if the Entity has free space. For Vertical and Grid types only. |

| [GridLayoutConstraint](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutConstraint) Constraint |
| --- |
| Constraints on the number of rows and columns. Grid type only. |

| int32 ConstraintCount |
| --- |
| The number of rows or columns to freeze based on the constraint. Grid type only. |

| [ChildAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/ChildAlignmentType) GridChildAlignment |
| --- |
| Sets the alignment direction of the child UI Entity if the Entity has free space. For Grid type only. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GridSpacing |
| --- |
| Horizontal and vertical spacing between child UI Entities. Grid type only. |

| [HorizontalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/HorizontalScrollBarDirection) HorizontalScrollBarDirection |
| --- |
| The orientation of the horizontal scrollbar. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it is placed in the world. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| Sets the free space on the top, bottom, left, and right of the layout group. |

| boolean ReverseArrangement |
| --- |
| Determines whether to reverse the existing sort. For Vertical and Horizontal types only. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarBackgroundColor |
| --- |
| The background color of the scrollbar. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarBgImageRUID |
| --- |
| The background image of the scrollbar. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) ScrollBarHandleColor |
| --- |
| The handle color of the scrollbar. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ScrollBarHandleImageRUID |
| --- |
| The handle image of the scrollbar. |

| float ScrollBarThickness |
| --- |
| The thickness of the scrollbar area. |

| [ScrollBarVisibility](https://mod-developers.nexon.com/apiReference/Enums/ScrollBarVisibility) ScrollBarVisible |
| --- |
| Sets whether to display the scrollbar automatically. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| float Spacing |
| --- |
| The spacing between child UI Entities. For Vertical and Horizontal types only. |

| [GridLayoutAxis](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutAxis) StartAxis |
| --- |
| The alignment direction of the child UI Entity. Grid type only. |

| [GridLayoutCorner](https://mod-developers.nexon.com/apiReference/Enums/GridLayoutCorner) StartCorner |
| --- |
| The starting position of the child UI Entity's alignment. Grid type only. |

| [LayoutGroupType](https://mod-developers.nexon.com/apiReference/Enums/LayoutGroupType) Type |
| --- |
| Sets the sort format of the layout group. |

| boolean UseScroll |
| --- |
| Sets whether to use the scroll function. |

| [VerticalScrollBarDirection](https://mod-developers.nexon.com/apiReference/Enums/VerticalScrollBarDirection) VerticalScrollBarDirection |
| --- |
| The orientation of the vertical scrollbar. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) GetScrollNormalizedPosition() |
| --- |
| Returns the normalized position of the scroll bar. |

| float GetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| Returns the normalized position of the specified directional scroll bar. |

| void ResetScrollPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis) |
| --- |
| Moves the position of the scroll bar on the specified axis to the first position. The initial position depends on the direction of the scroll bar. |

| void SetScrollNormalizedPosition([UITransformAxis](https://mod-developers.nexon.com/apiReference/Enums/UITransformAxis) axis, float value) |
| --- |
| Moves the position of the scroll bar on the specified axis to the designated normalized position. The direction of 0 and 1 is different depending on the direction of the scroll bar. |

| void SetScrollPositionByItemIndex(int32 index) |
| --- |
| Moves the scroll bar to a position where the child UI Entity at a specific index is visible. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [ScrollPositionChangedEvent](https://mod-developers.nexon.com/apiReference/Events/ScrollPositionChangedEvent) |
| --- |
| An event that occurs when the scroll position changes on a scrollable UI Entity. Only occurs if the UI Entity has ScrollLayoutGroupComponent or GridViewComponent. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of changing the CellSize of ScrollLayoutGroupComponent, fixing it to the first column, copying the ItemEntity up to InitItemCount, and adding it to ScrollLayoutGroupComponent.

Test by adding this component to an Entity with ScrollLayoutGroupComponent and setting the property value of ItemEntity and InitItemCount.

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

Enables side-scrolling movement and jump, as well as collision with RectTile. Movement is affected by gravity. Used if the current tile map mode is SideViewRectTileMap.

# Properties

| boolean ApplyClimbableRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, the character on the rotating or slanted ladder will be affected by the shape of the ladder. When false, the character will not be affected by the slant or rotation of the ladder. |

| float DownJumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Controls the speed of the bounce back up when jumping down. The greater the value, the higher the bounce. |

| boolean EnableDownJump ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Turns Jump Down on or off. |

| float JumpDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Controls the rate of the jump speed. The greater the value, the faster the fall. |

| float JumpSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Controls the speed of the bounce when jumping. The greater the value, the faster the jump. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) MoveVelocity ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Sets the movement speed. MovementComponent will be used to control the movement. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [RectTileInfo](https://mod-developers.nexon.com/apiReference/Misc/RectTileInfo) GetUnderfootTile() |
| --- |
| Checks the tile information currently being stepped on, and returns nil if there is no tile stepped on. Returns the left tile if it is above the boundary of two horizontally side-by-side tiles. |

| boolean IsOnGround() |
| --- |
| Checks the current contact with the ground. |

| void SetPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's position based on local coordinates. |

| void SetWorldPosition([Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) position) |
| --- |
| Sets an Entity's location to the world-based coordinates. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [RectTileCollisionBeginEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionBeginEvent) |
| --- |
| This event occurs when touching the collidable tile. |

| [RectTileCollisionEndEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileCollisionEndEvent) |
| --- |
| This event occurs when being freed from the collided tile. |

| [RectTileEnterEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileEnterEvent) |
| --- |
| This event occurs when entering a specific quadrangle tile. |

| [RectTileLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/RectTileLeaveEvent) |
| --- |
| This event occurs when leaving a specific quadrangle tile. |

# Examples

#### GetWallTile

This is an example of detecting the wall tile touching the side.

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
- [Controlling Character Movement from SideViewRectTileMap](/docs?postId=759)
- [Creating Maps with SideViewRectTile Mode](/docs?postId=758)

Update 2025-10-28 PM 02:21


# SkeletonGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the ability to draw and control skeleton resources created with Spine in the UI. Only resources created with Spine 4.1 can be used.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) AnimationNames |
| --- |
| Sets the animation for track 1. Animations play in index order. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Sets the skin. |

| boolean FlipX |
| --- |
| Determines whether to invert based on the X axis. |

| boolean FlipY |
| --- |
| Determines whether to invert based on the Y axis. |

| boolean Loop |
| --- |
| Indicates whether this will play on loop. |

| float PlayRate |
| --- |
| Sets the animation play speed. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveMode |
| --- |
| Sets how skeleton resource ratio, pivot, size, and so on are managed. |

| boolean RaycastTarget |
| --- |
| Becomes the subject of screen touch or mouse clicks if the value is set to true. The UI hidden behind will not receive screen touch and mouse click inputs. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SkeletonRUID |
| --- |
| Sets the RUID for the skeleton resource. |

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) SkinNames |
| --- |
| Sets the skin. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void AddAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Adds an animation to the track. Sets the track number, animation name, and misc properties to animationClip. Track 1 can't be used. |

| void AddEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Adds an empty animation to the track. Sets the track number and misc properties to animationClip. Track 1 can't be used. |

| void ClearTrack(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Empties the track. Track 1 can't be used. |

| [SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) GetCurrentAnimation(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the currently playing animation. |

| void SetAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Clear the track and add an animation. Specify the track number, animation name, and other properties in the animationClip. Track 1 is unavailable. |

| void SetAttachment([string](https://mod-developers.nexon.com/apiReference/Lua/string) slotName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attachmentName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Adds an attachment to the slot. If attachmentName is nil, the attachment is removed. |

| void SetEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Clear the track and add an empty animation. Specify the track number and other properties in the animationClip. Track 1 cannot be used. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [SkeletonAnimationCompleteEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationCompleteEvent) |
| --- |
| This event occurs when the skeleton animation is done playing. If it is set to play on a loop, the event will trigger each time it's completed. |

| [SkeletonAnimationEndEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationEndEvent) |
| --- |
| This event occurs when an animation ends while transitioning to another skeleton animation. |

| [SkeletonAnimationStartEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationStartEvent) |
| --- |
| This event occurs when an animation begins playing while a skeleton animation is being transitioned. |

| [SkeletonAnimationTimelineEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationTimelineEvent) |
| --- |
| This event occurs upon detecting an event that was registered to the animation timeline while the skeleton animation is playing. |

Update 2025-12-03 PM 05:12


# SkeletonRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the ability to draw and control skeleton resources created with Spine. Only resources created with Spine 4.1 can be used.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) AnimationNames ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the animation for track 1. Animations play in index order. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the color. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether to invert based on the X axis. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether to invert based on the Y axis. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to repeat playback of the animation. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the animation play speed. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SkeletonRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the RUID for the skeleton resource. |

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) SkinNames ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the skin. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void AddAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| Adds an animation to the track. Sets the track number, animation name, and misc properties to animationClip. Track 1 can't be used. |

| void AddEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| Adds an empty animation to the track. Sets the track number and misc properties to animationClip. Track 1 can't be used. |

| void ClearTrack(int32 trackIndex) |
| --- |
| Empties the track. Track 1 can't be used. |

| [SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) GetCurrentAnimation(int32 trackIndex) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the currently playing animation. |

| void SetAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| Clear the track and add an animation. Specify the track number, animation name, and other properties in the animationClip. Track 1 is unavailable. |

| void SetAttachment([string](https://mod-developers.nexon.com/apiReference/Lua/string) slotName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) attachmentName) |
| --- |
| Adds an attachment to the slot. If attachmentName is nil, the attachment is removed. |

| void SetEmptyAnimation([SkeletonAnimationClip](https://mod-developers.nexon.com/apiReference/Misc/SkeletonAnimationClip) animationClip) |
| --- |
| Clear the track and add an empty animation. Specify the track number and other properties in the animationClip. Track 1 cannot be used. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SkeletonAnimationCompleteEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationCompleteEvent) |
| --- |
| This event occurs when the skeleton animation is done playing. If it is set to play on a loop, the event will trigger each time it's completed. |

| [SkeletonAnimationEndEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationEndEvent) |
| --- |
| This event occurs when an animation ends while transitioning to another skeleton animation. |

| [SkeletonAnimationStartEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationStartEvent) |
| --- |
| This event occurs when an animation begins playing while a skeleton animation is being transitioned. |

| [SkeletonAnimationTimelineEvent](https://mod-developers.nexon.com/apiReference/Events/SkeletonAnimationTimelineEvent) |
| --- |
| This event occurs upon detecting an event that was registered to the animation timeline while the skeleton animation is playing. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

Update 2025-12-02 PM 01:55


# SliderComponent

Sets the values within the minimum and maximum ranges, and graphically displays the values.

# Properties

| [SliderDirection](https://mod-developers.nexon.com/apiReference/Enums/SliderDirection) Direction |
| --- |
| Sets the direction to display the minimum to maximum values graphically. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FillRectColor |
| --- |
| The color of the area to represent the value graphically. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) FillRectImageRUID |
| --- |
| The image RUID of the area to represent the value graphically. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) FillRectPadding |
| --- |
| Sets the free space of the area to display the value graphically. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) HandleAreaPadding |
| --- |
| The free space in the movable area with the slider handle. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) HandleColor |
| --- |
| The color of the slider handle. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) HandleImageRUID |
| --- |
| The image RUID of the slider handle. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) HandleSize |
| --- |
| The size of the slider handle. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it is placed in the world. |

| float MaxValue |
| --- |
| The maximum Value. |

| float MinValue |
| --- |
| The minimum Value. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| boolean UseHandle |
| --- |
| Sets whether to use the handle. |

| boolean UseIntegerValue |
| --- |
| Sets whether to use the value only as an integer. |

| float Value |
| --- |
| Current value. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SliderValueChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SliderValueChangedEvent) |
| --- |
| This event occurs when changing the Slider value. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This example sets the value range of SliderComponent to 0 - 100 while displaying the value in TextComponent.

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

Plays and manages sound.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) AudioClipRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the AudioClipRUID for audio clip playback. |

| boolean Bgm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to play as background music. |

| float HearingDistance ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the maximum distance from the listener entity where the sound can be heard. |

| boolean KeepBGM ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Seamlessly continue playing the audio clip if the previous background music is the same as the current background music. It is applied when the Bgm and PlayOnEnable properties are true and SoundComponent is enabled.<br>Background music does not play seamlessly when it is called using the PlayBGM() function of SoundService. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to repeat playback. |

| boolean Mute ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the mute state. |

| float Pitch ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the pitch and playback speed of the audio clip. The higher the value, the higher the pitch, and the faster the playback speed. Must set a value between 0 and 3. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to play an audio clip with Enable. |

| boolean SetCameraAsListener ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Adjusts the sound volume according to the distance between the center of the screen and the audio clip. When there is a listener Entity specified through the SetListenerEntity function, the Entity becomes the listener regardless of whether this property is active or not. |

| float Volume ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the volume. Must set a value between 0 and 1. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| float GetTimePosition() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the current play position of the audio clip in seconds. If the audio clip is not loaded, it logs the error as -1. |

| float GetTotalTime() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the total length of the audio clip in seconds. If the audio clip is not loaded, it returns -1 and logs the error. |

| boolean IsAudioClipLoaded() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns whether the audio clip corresponding to AudioClipRUID is loaded or not. GetTimePosition(), GetTotalTime(), SetTimePosition(timeInSecond, targetUserId) functions cannot be executed if no audio clip is loaded. |

| boolean IsPlaying([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Checks if music is being played or not. |

| boolean IsSyncedPlaying() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Checks whether the synced audio clip is playing. Returns true unless StopSyncedSound has been called separately after PlaySyncedSound was called, or the full duration of the audio clip has been played and stopped. |

| void Pause([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Pauses audio clip playback. |

| void Play([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Plays the audio clip. |

| void PlaySyncedSound() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| Plays the synchronized audio clip. If BGM property is true, it doesn't work. |

| void Resume([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Resumes playing the audio clip. |

| void SetListenerEntity([Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) entity, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Adjusts the volume by setting the listener Entity. The greater the distance between the SoundComponent and the listener Entity is, the quieter the sound. |

| void SetTimePosition(float timeInSecond, [string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Sets the current play position to the designated time period. If the audio clip is not loaded, the error is logged. |

| void Stop([string](https://mod-developers.nexon.com/apiReference/Lua/string) targetUserId = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Client&amp;color=violet) |
| --- |
| Stops playing the audio clip. |

| void StopSyncedSound() ![custom](https://img.shields.io/static/v1?label=&amp;message=Server&amp;color=palevioletred) |
| --- |
| Stops playing the synchronized audio clip. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [SoundPlayStateChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SoundPlayStateChangedEvent) |
| --- |
| Event that occurs when the playback state of the background music for the SoundService and sound effect for the SoundComponent change. |

# Examples

In this example, an audio clip (AudioClipRUID) is set up, and a local player is selected as the listener. The audio clip is then configured to become louder the closer the listener gets to the SoundComponent's location.

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
- [Changing Soundtrack](/docs?postId=117)
- [Creating Sound Effects](/docs?postId=578)

Update 2025-10-28 PM 02:21


# SpawnLocationComponent

The special Component only for the SpawnLocation Model.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# SpriteGUIRendererComponent

Outputs the sprite or animation clip on the UI.

# Properties

| [SpriteAnimClipPlayType](https://mod-developers.nexon.com/apiReference/Enums/SpriteAnimClipPlayType) AnimClipPlayType |
| --- |
| Sets the animation clip's playback type. You can choose to play only once, on repeat, etc. |

| int32 EndFrameIndex |
| --- |
| The last frame number when outputting animation. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it is placed in the world. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalPosition |
| --- |
| The image output location. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| The image size. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialId ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| float PlayRate |
| --- |
| The speed of animation playback. |

| [PreserveSpriteType](https://mod-developers.nexon.com/apiReference/Enums/PreserveSpriteType) PreserveSprite |
| --- |
| Defines how to preserve the original image's proportions/pivot/size, etc. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| int32 StartFrameIndex |
| --- |
| The starting frame number when outputting the animation. |

##### inherited from ImageComponent:

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Sets the default color for the image. |

| boolean DropShadow |
| --- |
| Sets whether to output shadows for the image. |

| float DropShadowAngle |
| --- |
| Sets the angle to output the shadow. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| Sets the shadow color. |

| float DropShadowDistance |
| --- |
| Distance between image and shadow. |

| float FillAmount |
| --- |
| Percentage of the image with Type setting Filled. Use values between 0 and 1. |

| boolean FillCenter |
| --- |
| Sets whether to fill the center of the image area when Type is set to Sliced or Tiled. |

| boolean FillClockWise |
| --- |
| Sets the direction of filling when FillMethod is set to Radial90, Radial180, or Radial360. It fills clockwise if the value is true. |

| [FillMethodType](https://mod-developers.nexon.com/apiReference/Enums/FillMethodType) FillMethod |
| --- |
| Filling method when setting the Filled type. |

| int32 FillOrigin |
| --- |
| You can set the starting point for filling when Type is set to Filled. If FillMethod is Horizontal or Vertical, you can use values from 0 to 1. If FillMethod is Radial90, Radial180, or Radial360, you can use values from 0 to 3. |

| boolean FlipX |
| --- |
| Determines whether to invert based on the X axis of an image. |

| boolean FlipY |
| --- |
| Determines whether to invert based on the Y axis of an image. |

| int32 FrameColumn ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| int32 FrameRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| int32 FrameRow ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is no longer used. Please use the AnimationClip Editor. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) ImageRUID |
| --- |
| The image RUID to be displayed on the screen. |

| boolean Outline |
| --- |
| Sets whether to output the image outline. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| Image outline color. |

| float OutlineWidth |
| --- |
| Outline thickness. |

| boolean RaycastTarget |
| --- |
| Becomes the subject of screen touch or mouse clicks if the value is set to true. The UI hidden behind will not receive screen touch and mouse click inputs. |

| [ImageType](https://mod-developers.nexon.com/apiReference/Enums/ImageType) Type |
| --- |
| How to display the image. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

##### inherited from ImageComponent:

| void SetNativeSize() |
| --- |
| Resizes the image to its original size. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

| [SpriteGUIAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerChangeFrameEvent) |
| --- |
| This event is raised when changing the sprite animation's frame. |

| [SpriteGUIAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerEndEvent) |
| --- |
| This event is raised when finishing the sprite animation's playback. |

| [SpriteGUIAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerEndFrameEvent) |
| --- |
| Initializes SpriteAnimPlayerEndEvent to the specified value. |

| [SpriteGUIAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerStartEvent) |
| --- |
| This event is raised when starting the sprite animation's playback. |

| [SpriteGUIAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteGUIAnimPlayerStartFrameEvent) |
| --- |
| This event is raised when playing the first frame of sprite animation. |

# Examples

This is an example of controlling sprite animation and using the Fill function. You can test by adding it to an Entity with SpriteGUIRendererComponent.

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
- [Basic UI Components](/docs?postId=744)
- [Creating UI](/docs?postId=64)

Update 2025-10-28 PM 02:21


# SpriteParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the ability to create a sprite particle effect.

# Properties

| boolean ApplySpriteColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether the Color property is applied to the Sprite to be used by the particle. Even if the property is false, the transparency value of Color is applied. |

| [SpriteParticleType](https://mod-developers.nexon.com/apiReference/Enums/SpriteParticleType) ParticleType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the type of particle to be created. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the SpriteRUID to be used as the particle. |

##### inherited from BaseParticleComponent:

| boolean AutoRandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| boolean Loop ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles repeatedly. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| float ParticleCount ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the number of particles. |

| float ParticleLifeTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle duration. |

| float ParticleSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle size. |

| float ParticleSpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set the play speed of particles. |

| boolean Prewarm ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined by the Sorting Layer. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from BaseParticleComponent:

| void Play() |
| --- |
| Play stopped particles. |

| void Stop() |
| --- |
| Stop playing particles. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

Applies the Color property of SpriteParticle to the new spriteRUID and Sprite. In addition, this example demonstrates controlling particle playback by tapping a designated key. To ensure proper functionality, you need to add the SpriteParticleComponent to the Entity.

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
- [Using Particles](/docs?postId=1036)
- [Utilizing Particles](/docs?postId=764)

Update 2025-12-02 PM 01:55


# SpriteRendererComponent

Outputs the sprite or animation clip.

# Properties

| [Dictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/Dictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is no longer used. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ClipName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Colors the Sprite. |

| [SpriteDrawMode](https://mod-developers.nexon.com/apiReference/Enums/SpriteDrawMode) DrawMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Sprite's draw mode. Simple, Sliced, and Tiled can be used. |

| int32 EndFrameIndex ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The last frame of the animation to be played. |

| boolean FlipX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether the sprite is inverted relative to the X axis. |

| boolean FlipY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines whether the sprite is inverted relative to the Y axis. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Doesn't perform automatic replace when designating the Map Layer name into SortingLayer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) MaterialID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the Id of the material to be applied to the renderer. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same Layer. A greater number indicates higher priority. |

| float PlayRate ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| For animation resources, you can specify the playback speed. It supports at least a value of 0 or more, and the higher the number, the faster it is. |

| RenderSettingType RenderSetting ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is no longer used. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The RUID of the sprite or animation clip. |

| int32 StartFrameIndex ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The starting frame of the animation to be played. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TiledSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the size of an area where the sprite will be drawn when DrawMode is Tiled, Sliced. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ChangeMaterial([string](https://mod-developers.nexon.com/apiReference/Lua/string) materialId) |
| --- |
| Replaces the material to be applied to the renderer. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [EmbededSpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeFrameEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerChangeFrameEvent. |

| [EmbededSpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerChangeStateEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerChangeStateEvent. |

| [EmbededSpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerEndEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerEndEvent. |

| [EmbededSpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/EmbededSpriteAnimPlayerStartEvent) |
| --- |
| This Event is deprecated. Use SpriteAnimPlayerStartEvent. |

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

| [SpriteAnimPlayerChangeFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeFrameEvent) |
| --- |
| This event is raised when changing the sprite animation's frame. |

| [SpriteAnimPlayerChangeStateEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerChangeStateEvent) |
| --- |
| This event is deprecated. |

| [SpriteAnimPlayerEndEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndEvent) |
| --- |
| This event is raised when finishing the sprite animation's playback. |

| [SpriteAnimPlayerEndFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerEndFrameEvent) |
| --- |
| This event is raised when playing the last frame of sprite animation. |

| [SpriteAnimPlayerStartEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartEvent) |
| --- |
| This event is raised when starting the sprite animation's playback. |

| [SpriteAnimPlayerStartFrameEvent](https://mod-developers.nexon.com/apiReference/Events/SpriteAnimPlayerStartFrameEvent) |
| --- |
| This event is raised when playing the first frame of sprite animation. |

# Examples

The following is an example of randomly determining the meso value of a meso item and applying a different Sprite according to the value range.

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
- [Making Animations](/docs?postId=595)
- [Sprite Color Adjustment](/docs?postId=116)

Update 2025-08-27 PM 04:56


# StateAnimationComponent

Specifies the animation to be played by state changes.

# Properties

| [SyncDictionary<string, string>](https://mod-developers.nexon.com/apiReference/Misc/SyncDictionary-2) ActionSheet ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The table the animation's name and AnimationClip are mapped to. Used when IsLegacy's value is true. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ReceiveStateChangeEvent(IEventSender sender, [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) stateEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Function which handles StateChangeEvent upon receiving. AnimationClipEvent is raised to play AnimationClip which is basically mapped to a State. |

| void RemoveActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key) |
| --- |
| Removes the element corresponding to key in StateToAvatarBodyActionSheet. Removes the element from ActionSheet if IsLegacy's value is true. |

| void SetActionSheet([string](https://mod-developers.nexon.com/apiReference/Lua/string) key, [string](https://mod-developers.nexon.com/apiReference/Lua/string) animationClipRuid) |
| --- |
| Adds elements to StateToAvatarBodyActionSheet. AvatarBodyActionStateName's property value of the AvatarBodyActionElement entity added as an element is animationClipRuid, and PlayerRate's property value is 1. Adds an element to ActionSheet when when IsLegacy's value is true. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) StateStringToAnimationKey([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| Returns the name of mapped Animation to the State. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [AnimationClipEvent](https://mod-developers.nexon.com/apiReference/Events/AnimationClipEvent) |
| --- |
| This event occurs when AnimationClip change is required. |

# Examples

This example makes a monster output a random hit animation.

You can test it by deleting the monster's basic StateAnimationComponent and adding the following Component, which is an extended StateAnimationComponent.

When attacking a monster after playing, you can see Orange Mushroom's various hit animations.

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

[override] string StateStringToAnimationKey (string stateName)
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

Provides a function to set the status behaviors and transition rules and control them using a custom StateType.

# Properties

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) CurrentStateName ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| You can get the name of the current state. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| boolean AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName, boolean reverseResult = false) |
| --- |
| Connects stateName's status and nextStateName's status. Returns false if it fails. When the return value of StateType's OnConditionCheck is true, the status transitions from stateName to nextStateName. If reverseResult's value is true, a status transition occurs when OnConditionCheck's return value is false. |

| boolean AddCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName, func -> boolean conditionCheckFunction, boolean reverseResult = false) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is now deprecated. Use another AddCondition (string, string, boolean) function. |

| boolean AddState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, func updateFunction = nil) ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This function is now deprecated. Use another AddState (string, Type) function. |

| boolean AddState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, Type stateType) |
| --- |
| Adds the status named stateName with a custom StateType. Returns false if it fails. |

| boolean ChangeState([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName) |
| --- |
| Forces the current state to be changed to the specified state. |

| void RemoveCondition([string](https://mod-developers.nexon.com/apiReference/Lua/string) stateName, [string](https://mod-developers.nexon.com/apiReference/Lua/string) nextStateName) |
| --- |
| Disconnects the state stateName from the state nextStateName. |

| void RemoveState([string](https://mod-developers.nexon.com/apiReference/Lua/string) name) |
| --- |
| Removes the specified state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [DeadEvent](https://mod-developers.nexon.com/apiReference/Events/DeadEvent) |
| --- |
| This event is fired when the target dies. Occurs when transitioning from State to DEAD state. |

| [ReviveEvent](https://mod-developers.nexon.com/apiReference/Events/ReviveEvent) |
| --- |
| This event is fired when the target is revived. |

| [StateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/StateChangeEvent) |
| --- |
| Occurs when changing state. |

# Examples

This component displays the current state of the entity through ChatBalloon.

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
- [Controlling Entity Status](/docs?postId=686)

Update 2025-08-27 PM 04:56


# StateStringToAvatarActionComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray)

This Component connects the appropriate Avatar Action State to the StateComponent's State to deliver it to the AvatarRendererComponent.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [BodyActionStateChangeEvent](https://mod-developers.nexon.com/apiReference/Events/BodyActionStateChangeEvent) |
| --- |
| This event occurs when the BodyAction state changes. |

Update 2025-08-27 PM 04:56


# StateStringToMonsterActionComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray)

This Component is deprecated. Use StateAnimationComponent.

# Properties

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [MonsterActionStateEvent](https://mod-developers.nexon.com/apiReference/Events/MonsterActionStateEvent) |
| --- |
| This is a deprecated event. Use StateAnimationComponent. |

Update 2025-08-27 PM 04:56


# TagComponent

You can find specific Entities by the tag added to the Entity. Use EntityService to find tags.

# Properties

| [SyncList<string>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Tags ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets tag list. Multiple tags can be set. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void AddTag([string](https://mod-developers.nexon.com/apiReference/Lua/string) tag) |
| --- |
| Adds tag value. |

| void RemoveTag([string](https://mod-developers.nexon.com/apiReference/Lua/string) tag) |
| --- |
| Deletes tag value. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

The following is an example of a game with the goal to reach the finish line within a set amount of time. The Qualified Tag is added when the character enters the area. The winner is displayed after 10 seconds by changing the player's NameTag with the qualified tag.

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
- [EntityService to Search for Entities](/docs?postId=201)

Update 2025-08-27 PM 04:56


# TextComponent

Displays the text on the screen. We recommend you use it in conjunction with UITransformComponent.

# Properties

| [TextAlignmentType](https://mod-developers.nexon.com/apiReference/Enums/TextAlignmentType) Alignment |
| --- |
| How to align text. |

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the Text property. |

| boolean BestFit |
| --- |
| Adjust the font size to fit the area size. |

| boolean Bold |
| --- |
| Sets whether to use bold text. |

| float ConstraintX |
| --- |
| Sets the maximum width to limit. Operates when sizefit is true. |

| float ConstraintY |
| --- |
| Sets the maximum height to limit. Operates when sizefit is true. |

| boolean DropShadow |
| --- |
| Sets whether to output the shadow of text. |

| float DropShadowAngle |
| --- |
| Sets the angle of the shadow. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) DropShadowColor |
| --- |
| The shadow color. |

| float DropShadowDistance |
| --- |
| Distance between the text and the shadow. |

| [FontType](https://mod-developers.nexon.com/apiReference/Enums/FontType) Font |
| --- |
| Font types. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor |
| --- |
| The color to use to render the text. |

| int32 FontSize |
| --- |
| Font size. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| If true, finds and displays the text matching the current language settings in LocaleDataSet. Text's property value is used as a key. The actual value of the text property does not change. |

| boolean IsRichText |
| --- |
| Sets whether to use rich text. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it is placed in the world. |

| float LineSpacing |
| --- |
| Adjusts line spacing. |

| int32 MaxSize |
| --- |
| Determines the maximum font size. |

| int32 MinSize |
| --- |
| Determines the minimum font size. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| The color of the text's outline. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) OutlineDistance ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This is a deprecated property. Use OutlineWidth. |

| float OutlineWidth |
| --- |
| Sets the width of the text outline. |

| [OverflowType](https://mod-developers.nexon.com/apiReference/Enums/OverflowType) Overflow |
| --- |
| Sets the handling method for when the text exceeds the horizontal area. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| Sets the margin of the text area. |

| boolean SizeFit |
| --- |
| Resizes the text area to fit the text. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text |
| --- |
| The content to be displayed. |

| boolean UseConstraintX |
| --- |
| Sets the width to be limited using the ConstraintX value. |

| boolean UseConstraintY |
| --- |
| Sets the height to be limited using the ConstraintY value. |

| boolean UseNBSP |
| --- |
| Specifies whether to enable line breaking at the character level. Line breaking is determined by characters, so the line breaks immediately when a character reaches the end of the line if the value is set to true. The line breaks by words to avoid incorrect line breaking if the value is set to false. |

| boolean UseOutLine |
| --- |
| Use the outline effect. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the text matching the current language settings in LocaleDataSet when IsLocalizationKey's value is true. A Key will be used as Text's property value. Returns Text's property value when IsLocalizationKey's value is false. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the height of the text area that is outputted to a fixed-width space. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the width of the input text area. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

# Examples

This is an example of printing text one character at a time.

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
- [Automatic Translation](/docs?postId=1072)
- [Basic UI Components](/docs?postId=744)
- [Understanding Localization](/docs?postId=951)

Update 2025-10-28 PM 02:21


# TextGUIRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Displays the text on the screen. We recommend you use it in conjunction with UITransformComponent.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the Text property. |

| boolean BestFit |
| --- |
| Adjust the font size to fit the area size. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomLeftColor |
| --- |
| This is the lower left color. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomRightColor |
| --- |
| This is the lower right color. |

| boolean ColorGradient |
| --- |
| Applies a color gradient to the text. Color gradients are multiplied by the text color. |

| float ConstraintX |
| --- |
| Sets the maximum width to limit. Operates when sizefit is true. |

| float ConstraintY |
| --- |
| Sets the maximum height to limit. Operates when sizefit is true. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Font |
| --- |
| Sets the font. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor |
| --- |
| The color to use to render the text. |

| float FontSize |
| --- |
| Font size. |

| [FontStyleType](https://mod-developers.nexon.com/apiReference/Enums/FontStyleType) FontStyle |
| --- |
| This is a style that's being applied to the text. |

| [GradientModes](https://mod-developers.nexon.com/apiReference/Enums/GradientModes) GradientMode |
| --- |
| Selects the gradient type to be applied. |

| [TextHorizontalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextHorizontalAlignmentOption) HorizontalAlignment |
| --- |
| This is the text's horizontal alignment. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| If true, finds and displays the text matching the current language settings in LocaleDataSet. Text's property value is used as a key. The actual value of the text property does not change. |

| boolean IsRichText |
| --- |
| Sets whether to use rich text. |

| float MaxSize |
| --- |
| Determines the maximum font size. |

| float MinSize |
| --- |
| Determines the minimum font size. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor |
| --- |
| Sets the outline color. |

| float OutlineWidth |
| --- |
| Sets the outline thickness. |

| [TextOverflowMode](https://mod-developers.nexon.com/apiReference/Enums/TextOverflowMode) Overflow |
| --- |
| Sets the handling method for when the text exceeds the horizontal area. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding |
| --- |
| Sets the margin of the text area. |

| int32 Page |
| --- |
| no description |

| boolean SizeFit |
| --- |
| Resizes the text area to fit the text. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| [TextRendererSpacingOption](https://mod-developers.nexon.com/apiReference/Misc/TextRendererSpacingOption) SpacingOption |
| --- |
| Sets the spacing options. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text |
| --- |
| The content to be displayed. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopLeftColor |
| --- |
| This is the upper left color. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopRightColor |
| --- |
| This is the upper right color. |

| boolean Underlay |
| --- |
| Sets whether a shadow will be applied. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) UnderlayColor |
| --- |
| Sets the shadow color. |

| float UnderlayOffsetX |
| --- |
| Sets the X axis location of the shadow. |

| float UnderlayOffsetY |
| --- |
| Sets the Y axis location of the shadow. |

| boolean UseConstraintX |
| --- |
| Sets the width to be limited using the ConstraintX value. |

| boolean UseConstraintY |
| --- |
| Sets the height to be limited using the ConstraintY value. |

| [TextVerticalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextVerticalAlignmentOption) VerticalAlignment |
| --- |
| This is the text's vertical alignment. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the text matching the current language settings in LocaleDataSet when IsLocalizationKey's value is true. A Key will be used as Text's property value. Returns Text's property value when IsLocalizationKey's value is false. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the height of the text area that is outputted to a fixed-width space. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the width of the input text area. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

Update 2025-12-02 PM 01:55


# TextGUIRendererInputComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

The content of this string is delivered to TextGUIRendererComponent.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the PlaceHolder property. |

| boolean AutoClear |
| --- |
| Sets whether to initialize the input area automatically upon completing text input. |

| int32 CharacterLimit |
| --- |
| The maximum number of characters that can be entered. |

| [InputContentType](https://mod-developers.nexon.com/apiReference/Enums/InputContentType) ContentType |
| --- |
| Specifies the type that can be inputted. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsFocused ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays on whatever is active at the time. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Finds the text matching the current language settings and displays as a default phrase in LocaleDataSet if the value is set to true. A Key will be used as PlaceHolder's property value. The actual value of PlaceHolder's property does not change. |

| [InputLineType](https://mod-developers.nexon.com/apiReference/Enums/InputLineType) LineType |
| --- |
| Specifies how to enter. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) PlaceHolder |
| --- |
| This is the default text that is shown when no text is entered. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) PlaceHolderColor |
| --- |
| Sets the PlaceHolder color. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The input content. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ActivateInputField() |
| --- |
| Enables text input. IsFoucsed value is changed to true a few frames after calling ActivateInputFied(). |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedPlaceHolder() |
| --- |
| Finds and returns the text matching the current language settings in LocaleDataSet when IsLocalizationKey's value is true. A Key will be used as PlaceHolder's property value. Returns PlaceHolder's property value when IsLocalizationKey's value is false. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

| [TextInputEndEditEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEditorEvent) |
| --- |
| The editor event occurs after changing the value of InputField. |

| [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent) |
| --- |
| The event occurs after changing the value of InputField. |

| [TextInputSubmitEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEditorEvent) |
| --- |
| An event that is called when the Enter key is pressed. |

| [TextInputSubmitEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEvent) |
| --- |
| An event that is called when the Enter key is pressed. |

| [TextInputValueChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEditorEvent) |
| --- |
| The editor event occurs when changing the value of InputField. |

| [TextInputValueChangeEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEvent) |
| --- |
| The event occurs when changing the value of InputField. |

# Examples

This is an example on entering an ID and password.

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
	self.IdInput.PlaceHolder = "Enter Id"
	       
	self.PasswordInput.Text = ""
	self.PasswordInput.PlaceHolder = "Enter Password"
	self.PasswordInput.ContentType = InputContentType.Password
}
  
Event Handler:
[entity: EntityPath]
HandleTextInputEndEditEvent (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	log("The entered Id is " .. self.InputId .. ".")
}
  
[entity: EntityPath]
HandleTextInputEndEditEvent2 (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	self.Password = text
	log("The entered password is " .. self.Password .. ".")
}
```

Update 2025-12-02 PM 01:55


# TextInputComponent

Receive a string and pass it to TextComponent.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the PlaceHolder property. |

| boolean AutoClear |
| --- |
| Sets whether to initialize the input area automatically upon completing text input. |

| int32 CharacterLimit |
| --- |
| The maximum number of characters that can be entered. |

| [InputContentType](https://mod-developers.nexon.com/apiReference/Enums/InputContentType) ContentType |
| --- |
| Specifies the type that can be inputted. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsFocused ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays on whatever is active at the time. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Finds the text matching the current language settings and displays as a default phrase in LocaleDataSet if the value is set to true. A Key will be used as PlaceHolder's property value. The actual value of PlaceHolder's property does not change. |

| boolean IsWorldUI ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether it is placed in the world. |

| [InputLineType](https://mod-developers.nexon.com/apiReference/Enums/InputLineType) LineType |
| --- |
| Specifies how to enter. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| boolean OverrideSorting ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Determines whether to set the SortingLayer and OrderInLayer values manually. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) PlaceHolder |
| --- |
| The default text displayed when no text is entered. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) PlaceHolderColor |
| --- |
| Sets the PlaceHolder color. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined according to the Sorting Layer. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The input content. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void ActivateInputField() |
| --- |
| Enables text input. IsFoucsed value is changed to true a few frames after calling ActivateInputFied(). |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedPlaceHolder() |
| --- |
| Finds and returns the text matching the current language settings in LocaleDataSet when IsLocalizationKey's value is true. A Key will be used as PlaceHolder's property value. Returns PlaceHolder's property value when IsLocalizationKey's value is false. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

| [TextInputEndEditEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEditorEvent) |
| --- |
| The editor event occurs after changing the value of InputField. |

| [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent) |
| --- |
| The event occurs after changing the value of InputField. |

| [TextInputSubmitEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEditorEvent) |
| --- |
| An event that is called when the Enter key is pressed. |

| [TextInputSubmitEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputSubmitEvent) |
| --- |
| An event that is called when the Enter key is pressed. |

| [TextInputValueChangeEditorEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEditorEvent) |
| --- |
| The editor event occurs when changing the value of InputField. |

| [TextInputValueChangeEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputValueChangeEvent) |
| --- |
| The event occurs when changing the value of InputField. |

# Examples

This is an example of receiving an input ID and password.

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
	self.IdInput.PlaceHolder = "Id"
	  
	self.PasswordInput.Text = ""
	self.PasswordInput.PlaceHolder = "password"
	self.PasswordInput.ContentType = InputContentType.Password
}
 
Event Handler:
[entity: EntityPath]
HandleTextInputEndEditEvent (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	log("The ID entered is " .. self.InputId .. ".")
}
 
[entity: EntityPath]
HandleTextInputEndEditEvent2 (TextInputEndEditEvent event)
{
	-- Parameters
	local text = event.text
	--------------------------------------------------------
	self.Password = text
	log("The password entered is " .. self.Password .. ".")
}
```

# SeeAlso

- [TextInputEndEditEvent](https://mod-developers.nexon.com/apiReference/Events/TextInputEndEditEvent)
- [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity)
- [Basic UI Components](/docs?postId=744)

Update 2025-10-28 PM 02:21


# TextRendererComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Displays text in the world space. Use it in conjunction with TransformComponent.

# Properties

| boolean AllowAutomaticTranslation ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Sets whether to automatically translate the value of the Text property. |

| boolean BestFit ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Adjusts the font size to fit the area size. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomLeftColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color on the bottom left. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) BottomRightColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color on the bottom right. |

| boolean ColorGradient ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Applies a gradient color to the letter. The letter color displayed will be multiplied by the gradient color applied. |

| float ConstraintX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the maximum width as the limit. Operates when sizefit is true. |

| float ConstraintY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the maximum height as the limit. Operates when sizefit is true. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Font ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the font. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) FontColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color in which to render the text. |

| float FontSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The font size. |

| [FontStyleType](https://mod-developers.nexon.com/apiReference/Enums/FontStyleType) FontStyle ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| This is a style that's being applied to the text. |

| GradientType GradientMode ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Selects the gradient type to be applied. |

| [TextHorizontalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextHorizontalAlignmentOption) HorizontalAlignment ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| How to align horizontal text. |

| boolean IgnoreMapLayerCheck ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Does not perform automatic substitution when the Map Layer's name is specified in SortingLayer. |

| boolean IsLocalizationKey ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| If true, finds and displays the text that matches the current language settings in LocaleDataSet. Text property value will be used as a key. The Text property itself does not change. |

| boolean IsRichText ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to use rich text. |

| float MaxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the maximum font size. |

| float MinSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the minimum font size. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same layer. A higher number indicates higher priority. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) OutlineColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the outline color. |

| float OutlineWidth ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the thickness of the border. |

| [TextOverflowMode](https://mod-developers.nexon.com/apiReference/Enums/TextOverflowMode) Overflow ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines what to do when the text exceeds the text box's width. |

| [RectOffset](https://mod-developers.nexon.com/apiReference/Misc/RectOffset) Padding ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the margins for the text area. |

| int32 Page ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Displays the text across multiple pages. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RectOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| A base point of the text box. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) RectSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The size of the text box. |

| boolean SizeFit ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Resizes the text box to fit the text. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When two or more entities overlap, the priority is determined by the Sorting Layer. |

| [TextRendererSpacingOption](https://mod-developers.nexon.com/apiReference/Misc/TextRendererSpacingOption) SpacingOption ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the spacing options. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) Text ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The content to be displayed. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopLeftColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color on the top left. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) TopRightColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The color on the top right. |

| boolean Underlay ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Creates a shadow behind the letter. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) UnderlayColor ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the shadow color. |

| float UnderlayOffsetX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the X axis location of the shadow. |

| float UnderlayOffsetY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the Y axis location of the shadow. |

| boolean UseConstraintX ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the width to be limited using the ConstraintX value. |

| boolean UseConstraintY ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the height limit using the ConstraintY value. |

| [TextVerticalAlignmentOption](https://mod-developers.nexon.com/apiReference/Enums/TextVerticalAlignmentOption) VerticalAlignment ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| How to align vertical text. |

| boolean Wrapping ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Creates a line break. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) GetLocalizedText() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Returns the text matching the current language settings in LocaleDataSet when IsLocalizationKey's value is true. A Key will be used as Text's property value. Returns Text's property value when IsLocalizationKey's value is false. |

| float GetPreferredHeight([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText, float width) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the height of the text area that appears in a fixed-width space. |

| float GetPreferredWidth([string](https://mod-developers.nexon.com/apiReference/Lua/string) preferredText) ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) ![custom](https://img.shields.io/static/v1?label=&amp;message=Yield&amp;color=saddlebrown) |
| --- |
| Gets the width of the input text area. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

| [SortingLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/SortingLayerChangedEvent) |
| --- |
| This event occurs when changing the component's SortingLayer. |

Update 2025-08-27 PM 04:56


# TileMapComponent

Provides terrain features in tile form. Only one tile map per map layer.

# Properties

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the color of the tile map. |

| boolean CreateFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Determines whether foothold is created. If false, foothold is not created. |

| float FootholdDrag ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Changes the player's friction on the tile map. The larger the value, the faster it decelerates. |

| float FootholdForce ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The force applied to the player above the tile map. If positive, the force is applied to the right, and if negative, to the left. |

| float FootholdWalkSpeedFactor ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Changes the player's speed on the tile map. The higher the value is, the faster the speed. |

| boolean IgnoreMapLayerCheck |
| --- |
| Does not perform automatic substitution when Map Layer name is specified in SortingLayer. |

| boolean IncludeFinishFoothold ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Determines whether to include the curved end of foothold when creating it. |

| boolean IsBlockVerticalLine |
| --- |
| Determines whether the Entity is blocked by vertical footholds in its tile map. Only pertains to Entities with RigidbodyComponent. |

| boolean IsOddGridPosition |
| --- |
| Positions the tile map off of the grid's base point. |

| int32 OrderInLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Determines the priority within the same Layer. A greater number indicates higher priority. |

| boolean PhysicsInteractable ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| If true, it may collide with a Dynamic rigid body (PhysicRigidbody) using the Physics feature. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SortingLayer ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| When two or more Entities overlap, the priority is determined according to the Sorting Layer. |

| TileMapVersion TileMapVersion ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The version of the tile map generation rule. This is the property for backward compatibility. |

| [DataRef](https://mod-developers.nexon.com/apiReference/Misc/DataRef) TileSetRUID |
| --- |
| Specifies the RUID of the tile set in the tile map. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [OrderInLayerChangedEvent](https://mod-developers.nexon.com/apiReference/Events/OrderInLayerChangedEvent) |
| --- |
| This event occurs when changing the Component's OrderInLayer. |

# Examples

The tile map's color changes if the character collides with the entity. It returns to the original color if they exit the collision range.

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

The Entity becomes touchable, so you can control its behavior when touching it.

# Properties

| boolean AutoFitOnce ![custom](https://img.shields.io/static/v1?label=&amp;message=MakerOnly&amp;color=salmon) |
| --- |
| Automatically changes TouchArea and Offset values based on the scale value of AvatarRendererComponent or SpriteRendererComponent when set to true. This only changes the very first time. |

| boolean AutoFitToSize |
| --- |
| If set to true, the TouchArea and Offset values ​​will automatically change whenever the Scale value of the AvatarRendererComponent or SpriteRendererComponent changes. |

| boolean DynamicTouchArea ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This property is deprecated. Use the AutoFitToSize property. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) Offset |
| --- |
| Sets the centerpoint of the touch area. |

| boolean RelayEventToBehind |
| --- |
| Sets whether or not to generate a touch event in the TouchReceiveComponent located behind in the rendering order. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) TouchArea |
| --- |
| The size of the touch area. The value of TouchArea may change in the middle when AutoFitToSize is true. |

| float TouchAreaUpdateTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) |
| --- |
| This is a deprecated property. Use the AutoFitToSize property. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [TouchEvent](https://mod-developers.nexon.com/apiReference/Events/TouchEvent) |
| --- |
| This event occurs when an entity is touched. The entity must have a TouchReceiveComponent for the event to occur. |

| [TouchHoldEvent](https://mod-developers.nexon.com/apiReference/Events/TouchHoldEvent) |
| --- |
| This event occurs when an entity is touched. The entity must have a TouchReceiveComponent for the event to occur. Cannot be triggered by a short touch. |

| [TouchReleaseEvent](https://mod-developers.nexon.com/apiReference/Events/TouchReleaseEvent) |
| --- |
| This event occurs when an entity is released from the touch. The entity must have a TouchReceiveComponent for the event to occur. Cannot be triggered by a short touch. |

# Examples

#### TouchEvent

This example demonstrates how to display a specific UI when touching an NPC. This kind of process is useful for various purposes like talking or trading with NPCs.

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

This is an example where the player touches the entity and moves it. TouchReceiveComponent must be added to an entity.

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

Indicates the location, rotation, and the size of an entity.

# Properties

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Position ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Shows coordinates relative to an Entity's parent. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) QuaternionRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Shows the rotation value of the Entity as a Quaternion. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Rotation |
| --- |
| Represents the rotation value of the Entity in Euler angles. In 2D, rotate using the Z axis. Synchronized by QuaternitonRotation. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) Scale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Represents the size ratio of the Entity. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) WorldPosition |
| --- |
| Represents the world-related coordinates of an Entity. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) WorldRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Represents an Entity's world-related rotation value in Euler angles. |

| float WorldZRotation |
| --- |
| Represents the value of the Z axis among the Euler angle rotation values based on the Entity's world. |

| float ZRotation |
| --- |
| Represents the Z-axis value among the Euler angle rotation values of the Entity. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| [FastVector3](https://mod-developers.nexon.com/apiReference/Misc/FastVector3) PositionAsFastVector3() |
| --- |
| Change the Position value to a FastVector3 type. |

| void Rotate(float angle) |
| --- |
| Rotate this Entity counterclockwise by angle. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToLocalDirection([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldDirection) |
| --- |
| Converts the direction entered from world coordinates to local coordinates. It is not affected by scale or position. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToLocalPoint([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) worldPoint) |
| --- |
| Converts world coordinates entered to local coordinates. It is affected by scale. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldDirection([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) localDirection) |
| --- |
| Converts the direction entered from local coordinates to world coordinates. It is not affected by scale or position. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) ToWorldPoint([Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) localPoint) |
| --- |
| Converts local coordinates entered to world coordinates. It is affected by scale. |

| void Translate(float deltaX, float deltaY) |
| --- |
| Move the coordinates of this Entity by deltaX, deltaY. |

| [FastVector3](https://mod-developers.nexon.com/apiReference/Misc/FastVector3) WorldPositionAsFastVector3() |
| --- |
| Change the WorldPosition value to a FastVector3 type. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

The following is an example of rotating an Entity at a constant speed using the ZRotation property.

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

If you write the code with the same operation using the Rotate function, it should be as follows:

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

The following example makes an entity free fall. Use the Translate function and delta to move the Entity by the current speed.

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
- [Changing Location, Size, and Rotation of Entities](/docs?postId=82)

Update 2025-08-27 PM 04:56


# TriggerComponent

Sets the entity's collision area on an entity and provides a function to detect a collision.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Available on previous systems where IsLegacy is true. Sets the center point position of the collider rectangle based on the Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) BoxSize ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Specifies the width and height of the rectangular collider. |

| float CircleRadius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The radius of the circular collider. Valid when ColliderType is Circle. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) ColliderName ![custom](https://img.shields.io/static/v1?label=&amp;message=Deprecated&amp;color=dimgray) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| This property is deprecated. Use CollisionGroup. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) ColliderOffset ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the center point position of the collider based on the Entity. Available on new systems with IsLegacy set as false. |

| [ColliderType](https://mod-developers.nexon.com/apiReference/Enums/ColliderType) ColliderType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the type of collider. Available on new systems with IsLegacy set as false. |

| [CollisionGroup](https://mod-developers.nexon.com/apiReference/Misc/CollisionGroup) CollisionGroup ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Set up a collision group. |

| boolean IsLegacy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets whether this Component will operate with the previous system. In a new system, the Collider is affected by TransformComponent's rotation and size. You can also use a circle-shaped collider by setting the ColliderType. |

| boolean IsPassive ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Will not perform any collision checking itself if the value is true. If TriggerComponents with true IsPassive collide, no event will occur. For an event to occur, at least one of the TriggerComponents' IsPassive must be set to false. You can improve the World performance by reducing unnecessary collision checking. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) PolygonPoints ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| The positions of the points that make up the polygonal colliding body. Valid when ColliderType is Polygon. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void OnEnterTriggerBody([TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent) enterEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is called when entering the trigger area. |

| void OnLeaveTriggerBody([TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent) leaveEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is called when exiting the trigger area. |

| void OnStayTriggerBody([TriggerStayEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerStayEvent) stayEvent) ![custom](https://img.shields.io/static/v1?label=&amp;message=ScriptOverridable&amp;color=blue) |
| --- |
| This function is called in every frame when entering and remaining in the trigger area. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [TriggerEnterEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerEnterEvent) |
| --- |
| This event occurts when TriggerComponent's area overlaps. |

| [TriggerLeaveEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerLeaveEvent) |
| --- |
| This event occurs when TriggerComponent's area no longer overlaps. |

| [TriggerStayEvent](https://mod-developers.nexon.com/apiReference/Events/TriggerStayEvent) |
| --- |
| This event occurs every frame when TriggerComponent's area overlaps. Please be aware that this may negatively affect your World's function. |

# Examples

The following is an example of creating an HP recovery space.

Increases HP while the player is inside the Entity's trigger area named Heal. You can use TriggerEnterEvent and TriggerLeaveEvent to know if the player is inside the trigger area.

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

This is the parent component of Tween types.

# Properties

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, this Component is removed from an Entity when the tween reaches its destination and ends or when Stop() is called directly. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the tween will play automatically upon starting the game. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The current playback state. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The starting position of the tween. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial rotation value of the tween. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial size of the tween. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets the playback subject. The default value is Default. You can control the tween with functions on both the server and the client. Server status is always synchronized to the client. Setting the execution area control to client only restricts the control of the tween to the client only and does not synchronize to the server.<br>Tween Control Functions: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The amount of elapsed time after starting the tween. The unit of measure is the second. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void Destroy() |
| --- |
| Changes CurrentState to Destroying and removes this Component. |

| void Pause() |
| --- |
| Pauses the tween. Changes CurrentState to Pausing. |

| void Play() |
| --- |
| Plays the tween. Changes CurrentState to Playing. |

| void RestartFromCurrentPosition() |
| --- |
| Restarts from current location. |

| void Stop(boolean reset) |
| --- |
| Stops the tween. Changes CurrentState to Idle. Specifying true for reset returns the tween's position and progress to the initial state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# TweenCircularComponent

Provides circular motion of an object.

# Properties

| float Degree ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Defines the total rotation angle. It rotates once if the value is 360. It rotates infinitely if the value is 0. |

| boolean IsClockwise ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Define the direction of rotation. Rotate clockwise if set to true. |

| boolean LookAtCenter ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| References the center point when rotating if the value is set to true. |

| float Radius ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Defines the turning radius. |

| float Speed ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Defines how many degrees per second to rotate. Rotates once per second if the value is 360. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, this Component is removed from an Entity when the tween reaches its destination and ends or when Stop() is called directly. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the tween will play automatically upon starting the game. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The current playback state. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The starting position of the tween. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial rotation value of the tween. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial size of the tween. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets the playback subject. The default value is Default. You can control the tween with functions on both the server and the client. Server status is always synchronized to the client. Setting the execution area control to client only restricts the control of the tween to the client only and does not synchronize to the server.<br>Tween Control Functions: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The amount of elapsed time after starting the tween. The unit of measure is the second. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| Changes CurrentState to Destroying and removes this Component. |

| void Pause() |
| --- |
| Pauses the tween. Changes CurrentState to Pausing. |

| void Play() |
| --- |
| Plays the tween. Changes CurrentState to Playing. |

| void RestartFromCurrentPosition() |
| --- |
| Restarts from current location. |

| void Stop(boolean reset) |
| --- |
| Stops the tween. Changes CurrentState to Idle. Specifying true for reset returns the tween's position and progress to the initial state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of making an Entity with `TweenCircularComponent` rotate infinitely.

```
local tween =self.Entity.TweenCircularComponent
tween.Degree = 0

-- Sets the rotation radius to 2.
tween.Radius = 2

-- Rotates at a speed of 180 degrees per second. It takes 2 seconds to complete one full rotation. 
tween.Speed = 180
```

# SeeAlso

- [Entity Section Movement](/docs?postId=122)

Update 2025-08-27 PM 04:56


# TweenFloatingComponent

Entity round trips up and down.

# Properties

| float Amplitude ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the distance between the lowest point to the highest point. The distance from the starting point to the lowest/highest point is 0.5 if the value is 1. |

| float OneCycleTime ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the duration from the lowest point to the highest point in seconds.<br>For example, duration of a rotation will be 2 seconds if the value is 1. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Defines the moving method. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, this Component is removed from an Entity when the tween reaches its destination and ends or when Stop() is called directly. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the tween will play automatically upon starting the game. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The current playback state. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The starting position of the tween. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial rotation value of the tween. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial size of the tween. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets the playback subject. The default value is Default. You can control the tween with functions on both the server and the client. Server status is always synchronized to the client. Setting the execution area control to client only restricts the control of the tween to the client only and does not synchronize to the server.<br>Tween Control Functions: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The amount of elapsed time after starting the tween. The unit of measure is the second. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| Changes CurrentState to Destroying and removes this Component. |

| void Pause() |
| --- |
| Pauses the tween. Changes CurrentState to Pausing. |

| void Play() |
| --- |
| Plays the tween. Changes CurrentState to Playing. |

| void RestartFromCurrentPosition() |
| --- |
| Restarts from current location. |

| void Stop(boolean reset) |
| --- |
| Stops the tween. Changes CurrentState to Idle. Specifying true for reset returns the tween's position and progress to the initial state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is an example of giving an Entity a floating effect.

```
local tween = self.Entity.TweenFloatingComponent
 
-- Movement distance 0.1
tween.Amplitude = 0.1
-- Frequency 1 second
tween.OneCycleTime = 1
-- Effect of gradually slowing and then getting faster again
tween.TweenType = EaseType.QuadEaseInOut
```

# SeeAlso

- [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType)

Update 2025-08-27 PM 04:56


# TweenLineComponent

Moves to the path designated by an entity. Supports various movement effects.

# Properties

| [CoordinateType](https://mod-developers.nexon.com/apiReference/Enums/CoordinateType) DestinationCoordinateType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether Positions is the absolute or relative coordinates. Restarts the tween if changed during playback. |

| float Duration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the duration from the departure point to the arrival point in seconds.<br>For example, if Duration is 1 for OneRoundTrip, each one-way trip duration will be 1 second in the forward direction and 1 second in the reverse direction. So the total duration will be 2 seconds. |

| [InterpolationType](https://mod-developers.nexon.com/apiReference/Enums/InterpolationType) Interpolation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the interpolation method for the movement path set in Positions. Restarts the tween if changed during playback. |

| [SyncList<Vector2>](https://mod-developers.nexon.com/apiReference/Misc/SyncList-1) Positions ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the movement path. Restarts the tween if changed during playback. |

| float ReturnDuration ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the duration returning from the destination to the starting point in seconds. Valid when UseReturnTweenType is true. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) ReturnTweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the movement effect returning from the destination to the starting point. Valid when UseReturnTweenType is true. |

| [TweenLinearStopType](https://mod-developers.nexon.com/apiReference/Enums/TweenLinearStopType) StopType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether it moves one way or round trips. |

| [EaseType](https://mod-developers.nexon.com/apiReference/Enums/EaseType) TweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets the workout type. |

| boolean UseReturnTweenType ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Sets whether to apply ReturnDuration and ReturnTweenType. Duration and TweenType will be applied upon returning from the destination to the starting point if the value is false. |

##### inherited from TweenBaseComponent:

| boolean AutoDestroy ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| When true, this Component is removed from an Entity when the tween reaches its destination and ends or when Stop() is called directly. |

| boolean AutoStart ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| If true, the tween will play automatically upon starting the game. |

| [TweenState](https://mod-developers.nexon.com/apiReference/Enums/TweenState) CurrentState ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The current playback state. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginPosition ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The starting position of the tween. |

| [Quaternion](https://mod-developers.nexon.com/apiReference/Misc/Quaternion) OriginRotation ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial rotation value of the tween. |

| [Vector3](https://mod-developers.nexon.com/apiReference/Misc/Vector3) OriginScale ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The initial size of the tween. |

| [TweenSyncType](https://mod-developers.nexon.com/apiReference/Enums/TweenSyncType) SyncType ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| Sets the playback subject. The default value is Default. You can control the tween with functions on both the server and the client. Server status is always synchronized to the client. Setting the execution area control to client only restricts the control of the tween to the client only and does not synchronize to the server.<br>Tween Control Functions: Play, Pause, Stop, Destroy, RestartFromCurrentPosition |

| number TweenTime ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| The amount of elapsed time after starting the tween. The unit of measure is the second. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from TweenBaseComponent:

| void Destroy() |
| --- |
| Changes CurrentState to Destroying and removes this Component. |

| void Pause() |
| --- |
| Pauses the tween. Changes CurrentState to Pausing. |

| void Play() |
| --- |
| Plays the tween. Changes CurrentState to Playing. |

| void RestartFromCurrentPosition() |
| --- |
| Restarts from current location. |

| void Stop(boolean reset) |
| --- |
| Stops the tween. Changes CurrentState to Idle. Specifying true for reset returns the tween's position and progress to the initial state. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Examples

This is example code which stops and executes Tween movement according to the key input.

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

Provides a function that can create particle effects with adjustable spawning areas in the UI.

# Properties

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaOffset |
| --- |
| Sets the location of the center point of the creation scope relative to the Entity. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) AreaSize |
| --- |
| Specifies the width and height of the particle spawning range. |

| [UIAreaParticleType](https://mod-developers.nexon.com/apiReference/Enums/UIAreaParticleType) ParticleType |
| --- |
| Sets the type of particle to be created. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| This indicates particle size. |

| boolean Loop |
| --- |
| Sets whether the particle will play on repeat. |

| float ParticleCount |
| --- |
| Set the number of particles. |

| float ParticleLifeTime |
| --- |
| Set the particle duration. |

| float ParticleSize |
| --- |
| Set the particle size. |

| float ParticleSpeed |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed |
| --- |
| Set the play speed of particles. |

| boolean Prewarm |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Play stopped particles. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Stops the particle from playing. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

Update 2025-08-27 PM 04:56


# UIBaseParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)![custom](https://img.shields.io/static/v1?label=&amp;message=Abstract&amp;color=darkkhaki)

The parent component of the UIParticleComponent, which creates the particle effect used in UI.

# Properties

| boolean AutoRandomSeed |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| This indicates particle size. |

| boolean Loop |
| --- |
| Sets whether the particle will play on repeat. |

| float ParticleCount |
| --- |
| Set the number of particles. |

| float ParticleLifeTime |
| --- |
| Set the particle duration. |

| float ParticleSize |
| --- |
| Set the particle size. |

| float ParticleSpeed |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed |
| --- |
| Set the play speed of particles. |

| boolean Prewarm |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Play stopped particles. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Stops the particle from playing. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# UIBasicParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides the function to set and control the default particles used in UI.

# Properties

| [UIBasicParticleType](https://mod-developers.nexon.com/apiReference/Enums/UIBasicParticleType) ParticleType |
| --- |
| Set the type of particle to be created. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| This indicates particle size. |

| boolean Loop |
| --- |
| Sets whether the particle will play on repeat. |

| float ParticleCount |
| --- |
| Set the number of particles. |

| float ParticleLifeTime |
| --- |
| Set the particle duration. |

| float ParticleSize |
| --- |
| Set the particle size. |

| float ParticleSpeed |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed |
| --- |
| Set the play speed of particles. |

| boolean Prewarm |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Play stopped particles. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Stops the particle from playing. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

Update 2025-08-27 PM 04:56


# UIGroupComponent

Indicates a UIGroup that groups UI Entities and sets the group's attributes. UIGroup can be created/deleted from the UIGroup edit window.

# Properties

| boolean DefaultShow |
| --- |
| Sets whether to activate the group by the start of the game. It will start as inactive when it's set to false. |

| int32 GroupOrder |
| --- |
| The layer order within the UIGroup. The higher GroupOrder's value is, the higher the layer is placed in the UIGroup edit window. |

| [UIGroupType](https://mod-developers.nexon.com/apiReference/Enums/UIGroupType) GroupType |
| --- |
| Sets the type of UIGroup.<br><br>* DefaultType: This is a default UIGroup. It's automatically created when the World is first created, and cannot be deleted.<br>* UIType: UIGroup created directly in the UI editor.<br>* EditorType: UIGroup created directly in the UI editor. It becomes the Editor UIGroup that can be used when editing the scene. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

Update 2025-08-27 PM 04:56


# UISpriteParticleComponent

![custom](https://img.shields.io/static/v1?label=&amp;message=Preview&amp;color=slategray)

Provides setting and control of basic particles.

# Properties

| boolean ApplySpriteColor |
| --- |
| Sets whether the Color property is applied to the Sprite to be used by the particle. Even if the property is false, the transparency value of Color is applied. |

| [UISpriteParticleType](https://mod-developers.nexon.com/apiReference/Enums/UISpriteParticleType) ParticleType |
| --- |
| Set the type of particle to be created. |

| [string](https://mod-developers.nexon.com/apiReference/Lua/string) SpriteRUID |
| --- |
| Sets the SpriteRUID to be used as the particle. |

##### inherited from UIBaseParticleComponent:

| boolean AutoRandomSeed |
| --- |
| Sets whether to create a new random seed whenever particle emission begins. |

| [Color](https://mod-developers.nexon.com/apiReference/Misc/Color) Color |
| --- |
| Corrects the color of the particles to be rendered. |

| boolean IsEmitting ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Indicates whether new particles are being emitted. |

| [Vector2](https://mod-developers.nexon.com/apiReference/Misc/Vector2) LocalScale |
| --- |
| This indicates particle size. |

| boolean Loop |
| --- |
| Sets whether the particle will play on repeat. |

| float ParticleCount |
| --- |
| Set the number of particles. |

| float ParticleLifeTime |
| --- |
| Set the particle duration. |

| float ParticleSize |
| --- |
| Set the particle size. |

| float ParticleSpeed |
| --- |
| Set the particle speed. |

| boolean PlayOnEnable |
| --- |
| Set whether to play particles when the particle component is set to Enable. |

| float PlaySpeed |
| --- |
| Set the play speed of particles. |

| boolean Prewarm |
| --- |
| If set to Enable, the maximum number of particles is loaded, and the particles play naturally. |

| integer RandomSeed |
| --- |
| Sets the random seed used to determine creation position, emission direction, speed, etc. when particle is played. |

##### inherited from Component:

| boolean Enable ![custom](https://img.shields.io/static/v1?label=&amp;message=Sync&amp;color=lightseagreen) |
| --- |
| Checks whether Component is activated or not. |

| boolean EnableInHierarchy ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) ![custom](https://img.shields.io/static/v1?label=&amp;message=HideFromInspector&amp;color=purple) |
| --- |
| Displays whether this Component is Enable in the hierarchy. If Entity Enable is false regardless of Component Enable, returns false. |

| [Entity](https://mod-developers.nexon.com/apiReference/Misc/Entity) Entity ![custom](https://img.shields.io/static/v1?label=&amp;message=ReadOnly&amp;color=orange) |
| --- |
| The Entity that owns this component. |

# Methods

##### inherited from UIBaseParticleComponent:

| void Play() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Play stopped particles. |

| void Stop() ![custom](https://img.shields.io/static/v1?label=&amp;message=ClientOnly&amp;color=orangered) |
| --- |
| Stops the particle from playing. |

##### inherited from Component:

| boolean IsClient() |
| --- |
| Returns whether the current execution environment is a client or not. |

| boolean IsServer() |
| --- |
| Returns whether the current execution environment is a server or not. |

# Events

| [ParticleEmitEndEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitEndEvent) |
| --- |
| This event is raised by BaseParticleComponent when emission of the particle has been completed. |

| [ParticleEmitStartEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleEmitStartEvent) |
| --- |
| The event that takes place when particle emission begins. |

| [ParticleLoopEvent](https://mod-developers.nexon.com/apiReference/Events/ParticleLoopEvent) |
| --- |
| If the Loop property is enabled, this event is fired when the particle's emission cycle returns and the emission repeats. |

Update 2025-08-27 PM 04:56

