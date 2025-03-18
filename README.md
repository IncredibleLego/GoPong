# GoPong
A simple pong implementation made in Go using Ebitengine

Started watiching this youtube video https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk

TODO:

CONSTANTS:

- Make screen adaptable to bigger dimensions
- Make menu adaptable to bigger dimensions (spacing between options based on screen dimension)

GAMEPLAY:

- Ball positioning at beginning (random)
- Ball delay after time hit
- Ball position at pause scene
- Ball bumping mechanic based on paddle hitpoint (angolo diverso in base a punto colpito diverso)
- 2 players and computer mode
- ball must go thowards player that just got scored

SCENE:

- Highscores saved to file
- Sounds

MENU:

- PAUSE menu that really pauses game and don't stop
- PAUSE menu with gamescene blurred in background
- Menu option to center in screen where you want
- STARTMENU with black background, selected option in menu flashes, default options are white
- STARTMENU: HIGHSCORES option to show the best highScores (ask player name and save?)
- New colors
- Import pictures?
- Center and manage menu
- Mouse selection full implementation

- OnEnter() OnExit in scenes: remove?

STAMPA VALORI:
due sprite 

serie di numeri

vuoto 1 2 3 4 5 6 7 8 9

al'inizio due sprite

vuoto 0
vuoto 1 ecc.
vuoto 9
1 0
1 1 ecc



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