# GoPong
A simple pong implementation made in Go using Ebitengine

Started watiching this youtube video https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk


- IMPLEMENT INPUT AS A GENERAL INPUT FUNCTION, ex
utils.Input(stringToEdit, maxCharacters)

- Create my own library with ebiten utils integration

## ROADMAP:

- When changing scale, the new scales must be rounded to the nearest 10 (done like optionscene.go row 55)

- popup.go in utils to generate confirm pop-ups ("confirm this decision?")

- In nameinputscene, if you keep a letter or backspace pressed you should be able to delete many characters or write more characthers

- Option "reset all values to default"
- Option to modify screen width and height together (ex 1080 x 640) etc, so uniting the two options in one

- Possibility to change commands? example choose the pause button?

- Make moveInterval inside menu an option calculated using menuoptions per second and remove that option

- Screen draw must be improved amd should print to the center of the screen
    - Maybe doing two functions: screenDraw with less arguments for all simple inputs and screenDraw extra for all different options

NEXT:
- Global variables such as: lastGameSceneId, lastSceneId, isInOptions to make easier menu managment

### GENERAL:
- Make game scalable: all that is drawn on screen must be adapted to the options
- Modify draw on screen options to make text (menu ecc.) related to textDimension 

### SAVES:

- Match save on file, saving data of a match on file to open it later
    -  When entering a mode ask if you wanna start a new match or restart from 0

- System to avoid ball always hitting in the same position (even if ball hits in the same point, it changes angle after a bit)
- Centre menu mouse options for bigger screens

- Color schemes in settings? (change color of menus as option)
- Menu position relative to number of options?
- Offsets to print not with simple numbers

- In AI mode, select difficulty after nameinputscene.go in the same file with a optionmenu (or simpler)

### OPTIONS:

- Options scene full implementation
- Set to default values options (for single values too, maybe a button to press)
- Implement mouse selection

### GAMEPLAY:

- Ball delay after time hit
- Modify increase speed, no longer working

### SCENES:

- Highscores saved to file
- Sounds
- Access time to scenes (don't start immediately but have some time)
- ()OnExit remove?

- AI MODE
    AI Features:
    - Doesn't always hit in the center, based on a random number generator it decides where to hit
    - Doesn't always hit the ball: difficulty based on a number of possibilities/hits

### MENU:

- ScreenDraw() method must be simpler with less passed values

- PAUSEMENU
    - Gamescene blurred in the background
    - Choose if the scene must be saved or not

- STARTMENU
    - Highscores option to show the best highScores (make player save name and save)
    - Unplayable pong playing in the background (https://richardcarter.org/)