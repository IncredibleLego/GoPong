# GoPong
A simple pong implementation made in Go using Ebitengine

Started watiching this youtube video https://youtu.be/V_OGeYj6p00?si=IWM1MB7iM3R7jqLk

TODO:

Mechanics improvement:

- Ball bugged at beginning

- Highscore saved to file
- Computer playing pong
- Sounds
- Random ball direction at start
- Speed increase
- Timer after ball hits

- Make collideWithPaddle a paddle method
- New color for menu

Menu a bottoni per scegliere tra le 3 modalità:
3 modalità:
- da solo
- con IA
- due giocatori

Differences:


- Solo: wall to hit
- IA: second paddle IA controlled
- Multi: second paddle, player controlled

Salva i risultati di highscore in un file ed estrai i risultati da visualizzare

es.

highScoreSolo: 5
highScoreIA = 4
highScoreMulti = 6

update in menu use of deprecated text.Draw function

UPDATE AND DRAW CORRECT MENU

POTER SELEZIONARE OPZIONI MENù CON MOUSE

constant textDimension to pass to all draw function to have default text value



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