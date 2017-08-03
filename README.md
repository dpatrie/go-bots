# go-bots

Client/Server Turn based code game

The goal of this game is to create Bots and make them compete against each other. To create a bot, you have to write code!

### Rules of the game
- A game has a 20x20 squares board. Each square can contain:
    - A bot
    - A power-up
    - A wall
    - Nothing
- Winner is the last bots standing or all bots still alive after 1000 turns.
- Each game can have 2 to 10 bots
- Each bot has 100 hit point
- Every turn, your bot receive information about nearby square (immediate) on the board
- Every turn, your bot can do one of theses actions
    - Look around (see more squares on next turn)
    - Defend (in case of attack, you take 1/2 hp and your attacker also get damage)
    - Move to a nearby square
    - Attack a nearby opponent (subtract between 10 to 20 hp, randomly, from your opponent)
    - Shoot a long range missiles. Direct hit -75 hp. Nearby square = -25 hp. (2 per game)
- Some squares can contain power-up's
    - Nitro boost (Allow to move 3 squares in one turn)
    - Missile crate (Gets 2 additional long range missiles)
    - Super vision (Can see the same thing as "Look around" for 2 turns)
    - Repair kit (Get back 25 hp, up to 100)
    - Super Shield (Invisible for 3 turn)
- Power-up's will appear randomly on the game board


#### Server Responsibilities
- Generate game board randomly
- Keeping the state of various games
    - Keep the state of board
    - Action History (to do a game replay)
    - Keep the state of each bot
        - Position
        - HP
        - Active Power up's
    - Order of play
- Broadcast message about the game to each bot in the said game
- Ask each bots for their move, giving them a context
- Provide a game viewer

#### Client Responsibilities
- Implement a Play() Method
- Respond to server play request in under 100ms

## Protocol
Client and server will talk a simple json protocol over TCP

### To list existing game
See existing game with remaining spot

**Client**

    {"request":"listGame"}

**Server reply (on success)**

    {
        "response":"ok", 
        "games":[
            {            
                "id" : 1,
                "name":"name of the game",
                "bots":[
                    "botName1",
                    "botName2",
                    ...
                ]
           },
           ...                    
        ]
    }

**Server reply (on error)**

    {"response":"error", "message":"error message"}

###To create a game
The game will automatically start after 60 seconds

**Client**

    {
        "request":"createGame",
        "param": {
            "name":"name of the game",
            "botName":"name of your bot",
            "width":30,
            "height":30
        }
    }
    
**Optional**

- width (default 20)
- height (default 20)

**Server reply (on success)**

    {"response":"ok", "gameId": 1}
    
**Server reply (on error)**

    {"response":"error", "message":"error message"}


### To join a game

**Client**

    {
        "request":"joinGame",
        "param": {
            "gameId":1,
            "botName": "name of your bot"
        }
    }

**Server reply (on success)**

    {"response":"ok"}

**Server reply (on error)**

    {"response":"error", "message":"error message"}

### To play a turn

Server will contact the client every turn with the context

**Server**

    {
        "request":"play",
        "param": {
            "bot" : {
                "name":"name of the bot",
                "hitPoint":100,
                "missileCount":2,
                "underAttack":false,
                "activePowerUp": ["nitroBoost", "superVision"],
                "position":{"x":2,"y":2}
            }
            "squares":[
                {
                    "position: {"x":2, "y":2},
                    "type":"empty|wall|powerup|bot",
                    "powerup":"nitroBoost|superVision|missileCrate|superShield",
                    "bot":{
                        "name":"name of the bot",
                        "hp":100,
                        "missileCount": 1,
                        "underAttack":false,
                        "activePowerUp": ["nitroBoost", "superVision"],                            
                    }
                },
                ...
            ]
        }
    }

**Client**

    {"response":"look|defend|move|attack|missile", "param" : {"x": 3, "y": 3}}

**Server reply (on success)**

    {"response":"ok"}

**Server reply (on error)**

    {"response":"error", "message":"error message"}
