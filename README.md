# GoPong
A Pong inspired Go version of the game, made using Ebitengine

First Inspiration: https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk

## ROADMAP:

- Highscores must be adjusted to always be in the center of the screen

### GENERAL:
- Make game scalable: all that is drawn on screen must be adapted to the options
- ALL OPERATIONS THAT CAN BE DONE ONCE, MUST BE DONE ONCE IN DRAW AND UPDATE
- Modify draw on screen options to make text (menu ecc.) related to textDimension 
- Global variables such as: lastGameSceneId, lastSceneId, isInOptions to make easier menu managment

### SAVES:
- Change values change in options after scale change (ex. Paddle width can't always change by 5)

### OPTIONS:
- Implement mouse selection in optionscene (not optionmenu)
- Color schemes in settings? (change color of menus as option). Color Presets
- Game presets: ex. slow ball fast paddle ecc.
- Possibility to change commands? example choose the pause button?
- Adjust max textDimension value (now 35)

### GAMEPLAY:
- Ball delay after time hit
- Modify increase speed, no longer working

### SCENES:
- ()OnExit remove?
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

- PAUSEMENU
    - Gamescene blurred in the background (like popups)

- STARTMENU
    - Unplayable pong playing in the background (https://richardcarter.org/)

### HIGHSCORES:
- 2 menus on top: one to choose gamemode (solo, multi, ai) and one to change the difficulty: so I have 1 page for all leaderboards
- Different leaderboards for different difficulties? Maybe to choose between an option above like

- Multiplayer highscores need adjustment -> When names are max possible it goes out of scale
- Multiplayer highscores reach errors like "16:60"

- Update and draw for the 3 will be in common except for the shown data and title
- Turn highscoreSelected in bool: menu or not menu?

- Highscores must be saved at the end of the match to avoid create many new records when the match is one
- Really highscores

- All highscores in one page? If yes how to change that

- Highscores shouldn't be saved at every point. Variable "last highscore" to save and save when you exit/game crashes

### POPUPS:
- Add mouse managment in popup
- Remove all magic numbers from popup, make it adapt to all options inserted

## PERFORMANCE:
- Create a static image for background (ex the net or the title) to avoid drawing it in every frame

### HARD/USELESS FEATURES

- Making unable to click on an option in the menu if the mouse isn't on it (the click is handled in the scene and not in the menu, it is hard to implement and pretty useless)
- Enabling mouse selection in optionmenu (really hard, should check if it's exactly on the arrows and increase/decrease according to that, could work but pretty hard)


INSERISCI UN EASTER EGG