i) Run the appcfg from within the appengine folder
ii) appcfg.py -A <YOUR_PROJECT_ID_> update appengine/
   did not work, but
   appcfg.py -A <YOUR_PROJECT_ID_> update engine did (without the trailing slash)

   cd to the higher level and running option
   appcfg.py -A <YOUR_PROJECT_ID_> update appengine/
   did not work either

 iii) I had to have a app.yaml for local "goapp serve"
 and a copy called "appengine.yaml" for the appcfg.py .
