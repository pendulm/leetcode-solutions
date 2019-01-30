cat words.txt |tr  " " "\n" |sed '/^$/d'|sort|uniq -c|sort -hr|awk '{print $2, $1}'
