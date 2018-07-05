#!/bin/bash
master_host=$1
USER=ec2-user

cat > ~/openShiftConfig.yml <<EOF
---
- hosts: ${master_host}
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

- hosts: ${master_host}
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