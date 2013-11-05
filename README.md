go-bots
=======

Client/Server Turn based code game

The goal of this game is to create Bots and make them compete against each other. To create a bot, you have to write code!

Rules of the game
- Winner is Last bots standing or bots still alive after 1000 turns.
- Order of play will randomize every 5 turns
- Each game can have 2 to 10 bots
- Each bot has 100 hit point
- Every turn, your bot receive information about nearby square (immediate) on the board
- Every turn, your bot can do one of theses actions
    - Look around (see more squares on next turn)
    - Defend yourself (in case of attack, you take 1/2 hp and your attacker also get damage)
    - Move to a nearby square
    - Attack a nearby opponent (subtract between 10 to 20 hp, randomly, from your opponent)
    - Shoot a long range missiles. Direct hit -75 hp. Nearby square = -25 hp. (2 per game)
- Some squares can contain power-up's
    - Nitro boost (Allow to move 3 squares in one turn)
    - Missile crate (Gets 2 additional long range missiles)
    - Super vision (Can see the same thing as "Look around" for 2 turns)
    - Repair kit (Get back 25 hp, up to 100)
- Power-up's will appear randomly on the game board


Server Responsabilities
- Generate game board randomly
    - A game is 20x20 square. Each square can contain:
        - A bot
        - A power-up
        - A wall

- Keeping the state of various games
    - Keep the state of board
    - Action History (to do a game replay)
    - Keep the state of each bot
        - Position
        - HP
        - Active Power up's
    - Order of play
- Publishing broadcast message about the game
- Publishing message to each bots


Client Responsabilities
- Implement a Play() Method
- Implement a events callback Method


Protocol
-----------
To be continued...
