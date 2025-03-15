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