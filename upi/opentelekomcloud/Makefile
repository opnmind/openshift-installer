
VERBOSE_LEVEL=-vvvv
VENV=~/.ansible_env
VENV_BIN=${VENV}/bin

up:
	. ${VENV_BIN}/activate && ansible-playbook -i inventory.yaml security-groups.yaml
	. ${VENV_BIN}/activate && ansible-playbook $(VERBOSE_LEVEL) -i inventory.yaml network.yaml

down:
	. ${VENV_BIN}/activate && ansible-playbook -i inventory.yaml down-security-groups.yaml
	. ${VENV_BIN}/activate && ansible-playbook -i inventory.yaml down-network.yaml

install: prepare fix

prepare:
	. ${VENV_BIN}/activate && ansible-galaxy collection install -r requirements.yaml --force

fix:
	cp /home/linux/security_group_rule.py /home/linux/.ansible/collections/ansible_collections/openstack/cloud/plugins/modules/
	cp /home/linux/openstack.py /home/linux/.ansible/collections/ansible_collections/openstack/cloud/plugins/module_utils/

venv:
	rm -rf ${VENV}
	virtualenv ${VENV}
	. ${VENV_BIN}/activate && pip install --upgrade pip
	. ${VENV_BIN}/activate && pip install -r requirements
	