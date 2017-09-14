### Настройка

В файле inventory надо указать ip адрес сервера, куда деплоим приложение, а так-же пользователя и ключ для авторизации:

```
ansible_ssh_private_key_file = ~/.ssh/id-rsa
ansible_user = root

[all]
10.0.0.1
```

Кол-во инстансов приложений указывается в переменной app_instance, которая находится в файле deploy.yml


```
---

- hosts: all
  become: yes
  become_method: sudo
  roles:
    - { role: common, app_user: app, app_group: app }
    - { role: deploy, app_user: app, app_group: app, app_instance: 5 }
```


### Деплой

```
ansible-playbook -i inventory deploy.yml
```

### Шаблоны
Шаблоны для генерации конфигов находятся в каталогах

roles/deploy/templates:

```
app.service  # systemd конфиг
config.json  # release конфиг
nginx.app    # nginx конфиг
```
