# Music for the blind
A tool for the blind to convert live music of **piano** into music sheet type MusicXML file using shortcut iphone, AnthemScore 4 and read by IBOS music reader

## Description
*  This method will use shortcut on iphone to automatically record live music of **piano** into audio and send it to **windows** computer to convert file audio to music XML (a file format performs music sheet in computer) and this XML file can be read by IBOS Music Reader

* This tool use [AnthemeScore 4](https://www.lunaverus.com/) to convert file audio of **piano** to MusicXML file
* File XML can be "read" by [IBOS Music read](https://www.ibos.dk/english/the-ibos-musicxml-reader.html) for the blind

## Set up

1. Download shortcut [Music Converter](https://www.icloud.com/shortcuts/12b0d94c0c854400a888863f75f0c1b8?fbclid=IwAR2kU8) for iphone
2. Download [AnthemeScore 4](https://www.lunaverus.com/) and
   [IBOS Music read](https://www.ibos.dk/english/the-ibos-musicxml-reader.html) for computer
3. Clone repository from github: 
 git clone https://github.com/tangducthinh456/Music-For-Blind
4. Open file main.go to config path to AnthemScore.exe, folder path to save audio file from iphone and folder path to save xml file
5. Turn on persional hotspot from iphone and let computer connect to Iphone's wifi
6. Open file main.go to config path to AnthemScore.exe, folder path to save audio file from iphone and folder path to save xml file <br/> Example <br/>
   const AudioSavePath = "C:\\Users\\Dell\\Desktop\\musictranscription\\Music-Transcription\\Audio\\"<br/>
const ExecuteProgramPath = "C:\\Program Files (x86)\\AnthemScore\\AnthemScore.exe"<br/>
const SaveXMLPath = "C:\\Users\\Dell\\Desktop\\musictranscription\\Music-Transcription\\MusicXML\\"
7. Get your local ip (use ipaddress on cmd)
8. Open Shortcut Music converter on iphone and change ip in URL block to local ip get from step 7
9. Open cmd, cd to folder Music-For-Blind and run<br/>
go build<br/>go run main.go
10. Turn on siri on iphone and start shortcut with "hey siri, Music converter", touch on screen and open or play the piano music
11. when finish the music, touch the iphone again and iphone will save file and send file to computer
12. Tool will receive file from iphone and use AnthemScore 4 to convert audio file to MusicXML file
13. Now you can open IBOS Music Reader to read music sheet created from live music in XML file
   

