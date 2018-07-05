#!/bin/bash
MASTER_HOST=$1
NAMESPACE=$2
OPUSER=$3
OPUSERPASS=$4
NODE1_HOST=$5
NODE2_HOST=$6
USER=$7

cat > ~/openShiftConfig.yml <<EOF
---
- hosts: ${MASTER_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  tasks:
    - htpasswd:
        path: /etc/origin/master/htpasswd
        name: $OPUSER
        password: '$OPUSERPASS'
        owner: root
        group: root
        mode: 0600

- hosts: ${MASTER_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  become_user: $USER
  tasks:
    - name: create a new project
      command: "{{item}}"
      with_items:
      - "oc login -u $OPUSER -p $OPUSERPASS"
      - "oc new-project $NAMESPACE"
      - "oc project $NAMESPACE"
      
- hosts: ${MASTER_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  become_user: $USER
  tasks:
    - name: create and configure service account
      command: "{{item}}"
      with_items:
      - "oc login -u system:admin -n $NAMESPACE"
      - "oc create sa qlik -n $NAMESPACE"
      - "oc adm policy add-scc-to-user anyuid system:serviceaccount:$NAMESPACE:qlik"
      - "oc adm policy add-scc-to-user hostaccess system:serviceaccount:$NAMESPACE:qlik"
      - "oc adm policy add-scc-to-user hostmount-anyuid system:serviceaccount:$NAMESPACE:qlik"
      - "oc adm policy add-cluster-role-to-user admin system:serviceaccount:$NAMESPACE:qlik -n $NAMESPACE"
      - "oc policy add-role-to-group edit system:serviceaccounts -n $NAMESPACE"
      - "oc label node ${NODE1_HOST} purpose=message"
      - "oc label node ${NODE2_HOST} purpose=db"

- hosts: ${NODE2_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  tasks:
    - name: create postgresql data directory
      file:
        path: /var/pgsql
        state: directory

    - name: Set postgresql data directory permissions
      file:
        path: /var/pgsql
        state: directory
        owner: 26
        group: 26
        mode: 0700
        setype: svirt_sandbox_file_t

- hosts: ${MASTER_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  become_user: $USER
  tasks:
    - name: download postgresql db
      get_url:
        url: https://github.com/xqianwang/micro-message/releases/download/1.1.1/database-1.1.1.tar.gz
        dest: ~/database.tar.gz
    
    - name: create postgresql directory
      file:
        path: ~/postgresql
        state: directory

    - name: untar crunchy db
      unarchive:
        remote_src: yes
        src: ~/database.tar.gz
        dest: ~/postgresql

    - name: deploy database app
      command: "{{item}}"
      args:
        chdir: ~/postgresql/
      with_items:
      - "oc login -u system:admin -n $NAMESPACE"
      - "~/postgresql/deploy.sh"    

- hosts: ${MASTER_HOST}
  remote_user: $USER
  become: yes
  become_method: sudo
  become_user: $USER
  tasks:
    - name: download micro-message app openshift files
      get_url:
        url: https://github.com/xqianwang/micro-message/releases/download/1.1.1/micro-message-1.1.1.tar.gz
        dest: ~/micro-message.tar.gz

    - name: create zrpe directory
      file:
        path: ~/micro-message
        state: directory

    - name: untar micro-message
      unarchive:
        remote_src: yes
        src: ~/micro-message.tar.gz
        dest: ~/micro-message

    - name: deploy micro-message application
      command: "{{item}}"
      args:
        chdir: ~/micro-message/
      with_items:
      - "oc login -u system:admin -n $NAMESPACE"
      - "~/micro-message/deploy.sh"
