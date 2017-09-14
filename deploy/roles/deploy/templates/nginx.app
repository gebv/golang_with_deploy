server {
  listen 80;
  location / {
     proxy_pass http://appcluster;
     proxy_pass_header x-mobile-device-id;
  }

}
    
upstream appcluster {
{% for item in ports %}
   server 127.0.0.1:{{ item }};
{% endfor %}
}
