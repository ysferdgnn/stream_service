# Stream Service
A basic media stream api made with Go language.Streams video,text,music and image files as playable with browsers.
Reads requested files from file system. Calculate piece counts by given chunk size and writes response bytes chunk by chunk.

# Media Types
  - video : mp4
  - music : mp3
  - text  : txt
  - image : jpg
# Files
 api search request file name under '/files' path and begin stream file from path.
# Routes
 - Video

    Searches video file with request name.
   ```
   GET /video/{name} 
   ```
  - Music

    Searches music file with request name.
    ```
    GET /music/{name}
    ```
   - Text
     
     Searches text file with request name.
     ```
     GET /text/{name}
     ```
   - Image
     
     Searches image file with request name.
     ```
     GET /image/{name}
     ```
 # Test Routes
  - GET   /video/test.mp4
  - GET   /video/test
  
  - GET   /music/test.mp3
  - GET   /music/test
  
  - GET   /text/test.txt
  - GET   /text/test
  
  - GET   /image/test.jpg
  - GET   /image/test
  

   
   
