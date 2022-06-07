# JourneyDiary2MD

**Journey.cloud diary's JSON to Markdown**

The project aims to make easy the transition from the online journal to a more **Private** and **Local** one

I use both Obsidian and Logseq for various logging purposes so I thought of converting everything in **Markdown** with the use of tags to enhance the searchability 

To export your notes, you can go in your app settings --> Data --> export/backup --> (choose the date range you prefer) --> done!

-> Extract the zip and you're ready to begin! 

---

To execute it:

- check the configuration file `config.json`
    ```json
    {
      "scanFolder" : "../journey/",   // folder containing all the jsons and assets from the exported zip
      "destFolder" : "../journals/",  // where to put generated .md files
      "assetsFolder" : "../assets/",  // where to put assets
      "fileNameFormat" : "2006_01_02",// Markdown files name format
      "isLogseq" : true,              // Do you want to have a bullet list indentation?
      "splitSameDayNotes" : false     // Do you prefer to have only one file for every day or you prefer to have multiple file_1, file_2 files?
    }
    ```
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
- [ ] safer error handling
- [ ] find out what the unused keys do
