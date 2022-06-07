# JourneyDiary2MD

Journey.cloud diary's JSON to Markdown

At the current state the procedure creates 2 folders
- assets
- journals

It assumes to have a folder structure like this:

+ go        // a folder that contains the repository code
  - main.go
+ journey   // the exported notes extracted (json + images + audio)
  - sadbhsbafhas.json
  - fbafbgafassa.json
  - ...
+ assets
+ journals

---

To execute it:
- respect the correct folder structure
- run ```go run .``` in the folder where you put the main.go

---

TODO
- [x] parse text
- [x] parse weather
- [x] parse location
- [x] html2md
- [x] parse assets
- [x] merge same-day notes in a single file
- [x] add configuration parameters to choose folders
- [ ] add error handling
- [ ] find out what the unused keys do
