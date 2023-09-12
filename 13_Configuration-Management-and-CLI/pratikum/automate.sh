#!/bin/bash

# ambil input command line argument 1
folder_utama="$1 $(date)"
mkdir "$folder_utama" && cd "$folder_utama"

# ambil input command line argument 2 dan 3
facebook_content="https://www.facebook.com/$2"
linkedin_content="https://www.linkedin.com/in/$3"

mkdir -p about_me/personal/ && mkdir about_me/professional/
echo $facebook_content > about_me/personal/facebook.txt
echo $linkedin_content > about_me/professional/linkedin.txt

# simpan curl ke dalam variable
myfriends_content=$(curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt)
mkdir my_friends && echo -e $myfriends_content > my_friends/list_of_my_friends.txt

# system info
mkdir my_system_info
about_this_content="My username: $(uname -n) \nWith host: $(uname -srvmpio)"
echo -e $about_this_content > my_system_info/about_this_laptop.txt

# ping
ping_info=$(sudo ping -c 3 google.com)
echo -e $ping_info > my_system_info/internet_connection.txt



