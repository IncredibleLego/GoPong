# GoPong
A Pong inspired Go version of the game, made using Ebitengine

First Inspiration: https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk

## ROADMAP:

- Add mouse managment in popup
- Remove all magic numbers from popup, make it adapt to all options inserted

### GENERAL:
- Make game scalable: all that is drawn on screen must be adapted to the options
- Modify draw on screen options to make text (menu ecc.) related to textDimension 
- Screen draw must be improved amd should print to the center of the screen
    - Maybe doing two functions: screenDraw with less arguments for all simple inputs and screenDraw extra for all different options
- Global variables such as: lastGameSceneId, lastSceneId, isInOptions to make easier menu managment
- Try implementing other scales to see if that is working

### SAVES:
- Match save on file, saving data of a match on file to open it later?  --> idk if it might make sense
    -  When entering a mode ask if you wanna start a new match or restart from 0
- Change values change in options after scale change (ex. Paddle width can't always change by 5)

### OPTIONS:
- Implement mouse selection
- Color schemes in settings? (change color of menus as option). Color Presets
- Game presets: ex. slow ball fast paddle ecc.
- Possibility to change commands? example choose the pause button?

- Don't restart immediately when selecting the new scale
    - Let user select the preferred scale
    - When exiting --> Popup message asking to reload (or just reload)
    - This makes users able to jump from a scale to one not immediately after it

### GAMEPLAY:
- Ball delay after time hit
- Modify increase speed, no longer working

### SCENES:
- ()OnExit remove?
- Name positioning in game scenes
- Turn nameimputscene in gameset scene where you select name, mode and difficulty

- AI MODE
    AI Features:
    - Doesn't always hit in the center, based on a random number generator it decides where to hit
    - Doesn't always hit the ball: difficulty based on a number of possibilities/hit
- In AI mode, select difficulty after nameinputscene.go in the same file with a optionmenu (or simpler)

### MENU:
- ScreenDraw() method must be simpler with less passed values
- Change menu and positioning after resizing
- Menu position relative to number of options?
- Make moveInterval inside menu an option calculated using menuoptions per second and remove that option

- PAUSEMENU
    - Gamescene blurred in the background (like popups)
    - Choose if the scene must be saved or not

- STARTMENU
    - Unplayable pong playing in the background (https://richardcarter.org/)

### HILIGHTS:
- Saved to file like options
- 2 menus on top: one to choose gamemode (solo, multi, ai) and one to change the difficulty: so I have 1 page for all leaderboards
- Hilights show date/time, name and points. For mode against AI is also shown difficulty
- Different leaderboards for different difficulties? Maybe to choose between an option above like

## PERFORMANCE:
- Create a static image for background (ex the net or the title) to avoid drawing it in every frame