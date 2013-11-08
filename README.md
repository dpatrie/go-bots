#go-bots
---

Client/Server Turn based code game

The goal of this game is to create Bots and make them compete against each other. To create a bot, you have to write code!

###Rules of the game
- A game is 20x20 square. Each square can contain:
    - A bot
    - A power-up
    - A wall
    - Nothing
- Winner is the last bots standing or all bots still alive after 1000 turns.
- Order of play will randomize every 5 turns
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
- Power-up's will appear randomly on the game board


####Server Responsibilities
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

####Client Responsibilities
- Implement a Play() Method
- Respond to server play request in under 100ms

##Protocol
Client and server will talk a simple json protocol over TCP

###To create a game

**Client**

`{"request":"create_game", "name":"name of the game", "password":"password to join game"}`

**Server reply**

`{"response":"created|error", "error":"error message if any", "game_id":"id of game", "client_id":"id of client"}`

###To join a game

**Client**

`{"request":"join_game", "game_id":"id of game to join"}`

**Server reply**

`{"response":"joined|error", "error":"error message if any", "client_id":"id of client"}`

###To play a turn

**Server**

    {
        "request":"play", 
        "client_id":"id of client", 
        "context":{
            "hp":100,
            "missile":2,
            "under_attack":true,
            "position":[2,2], 
            "square":[
                {
                    "position":[1,1],
                    "type":"empty|wall|powerup|bot",
                    "powerup":{},
                    "bot":{
                        "client_id":"",
                        "name":"",
                        "hp":100,
                        "under_attack":false
                    }
                },          
                {
                    "position":[1,2],
                    "type":"empty|wall|powerup|bot",
                    "powerup":{type:"nitro|vision|missile|repair"},
                    "bot":{}
                },                
                ...
            ]
        }
    }

**Client**

    {
        "response":"look|defend|move|attack|missile",
        "position":[3,3]
    }
