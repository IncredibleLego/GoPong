# GoPong
A simple pong implementation made in Go using Ebitengine

Started watiching this youtube video https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk

TODO:


CONSTANTS:

- Make game scalable: all that is drawn on screen must be adapted to the options
- Modify draw on screen options to make text (menu ecc.) related to textDimension

OPTIONS:

- Options saved to file, user adaptable
- Predefined options saved to restore original value (maybe original value for single values too?)

GAMEPLAY:

- Ball positioning at beginning (random)
- Ball delay after time hit

- Modify increase speed, no longer working

SCENES:

- Highscores saved to file
- Sounds
- Access time to scenes (don't start immediately but have some time)

SCENES:

- ()OnExit remove?
- Print score values: 2 sprite for all values (max value 99) with all numbers that increase
- When you exit a mode, it shouldn't be saved the score and results

- AI MODE
    - Beatable AI

MENU:

- ScreenDraw() method must be simpler with less passed values

- PAUSEMENU
    - Game paused that resmes after pause
    - Gamescene blurred in the background

- STARTMENU
    - Highscores option to show the best highScores (make player save name and save)
    - Unplayable pong playing in the background (https://richardcarter.org/)


Print colors:  
Colori Primari e Secondari  
Colore	Scala (r, g, b, a)  
Bianco	(1, 1, 1, 1)  
Nero	(0, 0, 0, 1)  
Rosso	(1, 0, 0, 1)  
Verde	(0, 1, 0, 1)  
Blu	(0, 0, 1, 1)  
Giallo	(1, 1, 0, 1)  
Ciano	(0, 1, 1, 1)  
Magenta	(1, 0, 1, 1)  
Colori Pastello e Tonalità Intermedie  
Colore	Scala (r, g, b, a)  
Grigio chiaro	(0.8, 0.8, 0.8, 1)  
Grigio scuro	(0.2, 0.2, 0.2, 1)  
Arancione	(1, 0.5, 0, 1)  
Rosa	(1, 0.5, 0.7, 1)  
Lime	(0.5, 1, 0, 1)  
Azzurro	(0.3, 0.6, 1, 1)  
Viola	(0.6, 0, 1, 1)  
Marrone	(0.6, 0.3, 0, 1)  
Colori Scuri  
Colore	Scala (r, g, b, a)  
Rosso scuro	(0.5, 0, 0, 1)  
Verde scuro	(0, 0.5, 0, 1)  
Blu scuro	(0, 0, 0.5, 1)  
Viola scuro	(0.4, 0, 0.6, 1)  