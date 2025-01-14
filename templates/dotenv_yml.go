package templates

const DOTENV_YML = `
---
- name: 'Trellis CLI: Template .env files to local system'
  hosts: web:&{{ env }}
  connection: local
  gather_facts: false
  tasks:
    - name: Template .env files to local system
      template:
        src: roles/deploy/templates/env.j2
        dest: "{{ item.value.local_path }}/.env"
        mode: '0644'
      with_dict: "{{ wordpress_sites }}"
`
