# Template
facebook.com/matkinhig
matkinhig@gmail.com
matkinhig@outlook.com

## INSTALL PROJECT WITH DOCKER COMPOSE.

B1 : AutoMigrate data hoặc run script init.sql cùng folder <br/>

B2 : init go tool hoặc cài đặt golang <br/>

B3 : config docker-compose.yaml <br/>

B4 : run cript : docker-compose up <br/>

B5 : test service : docker ps -a <br/>
<br/>
NẾU MUỐN KẾT NỐI VỚI DataBase ĐANG CHẠY BỞI DOCKER : docker exec -it "namesOfDataBase" -l <br/>

mysql -u root -p ronglong01 <br/>

## INSTALL ELASTICSEARCH
1. Install container Elastic Search
2. Install container Kibana
3. run script : docker-compose up

Yet, over time I found I was making the same changes to Bones at the start of every project and this was taking up valuable time. Thus, I created the predecessor to this theme called 'Osseous' which included some of the changes found in Template. Osseous literally means: 'of, relating to, or composed of bone' and was a good departure point.

Template picks up where Osseous left off and takes things a bit further with 'Template'. It includes namespaced functions and customized defaults along with a lot of other stuff I like and use for my development. Your mileage may vary.

2017 Update: we've gone through and done a major rewrite, adding and updating the following:
- WordPress Customizer support
- WooCommerce support
- updated body class function
- expanded Quicktags
- template part library (really cool)
- updated comments
- cleaned up header.php
- admin and login page updates
- updated media query .scss stylesheets
- removed @2x and ie_grid (who uses those?)
- HTML schema support
- updated css reset
- default .scss classes

...and much, much more.

This is the beginning for Template so we welcome improvements, comments, criticism and general feedback. I've probably made a ton of mistakes so as with anything open source, it is a work in progress.

## Recommended Plugins
Some of the stuff in Template references plugins that I use with just about every site and I recommend:
- Advanced Custom Fields (Pro): https://www.advancedcustomfields.com. A must-have.
- WP Retina 2x: https://wordpress.org/plugins/wp-retina-2x/. Works with WP built-in retina support. Just install and forget it.
- EWWW Image Optimizer: https://wordpress.org/plugins/ewww-image-optimizer/. Smush all of the images. Automagically.
- Plugin Organizer: https://wordpress.org/plugins/plugin-organizer/. Selective plugin loading on a per-page, per-post and per-type basis. Word.
- WP Migrate DB Pro: https://deliciousbrains.com/wp-migrate-db-pro/. Indispensable tool if you work on sites locally (and you should).


## Other stuffs
Designed by Joshua Michaels for studio.bio: http://studio.bio/themes/template

With help from @joniler.

License: WTFPL
License URI: http://sam.zoy.org/wtfpl/

Do whatever you want. Freedom, baby.

#### Special Thanks to:
@eddiemachado — all credit is due to him and the original Bones collaborators: Paul Irish, Yoast, Andrew Rogers, David Dellanave and others.


#### Submit Bugs & or Fixes:
https://github.com/joshuaiz/template/issues

To view Release & Update Notes, read the CHANGELOG.md file in the main folder.







cài đặt ElacticSearch bằng docker



