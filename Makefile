VENV := $(CURDIR)/.env

PYTHON := python3.11
PIP := $(VENV)/bin/pip
ANSIBLE_GALAXY := $(VENV)/bin/ansible-galaxy
ANSIBLE_PLAYBOOK := $(VENV)/bin/ansible-playbook

ANSIBLE_DIR := $(CURDIR)/ansible
INVENTORY := $(ANSIBLE_DIR)/inventory/hosts.yml

.PHONY: venv argo clean

venv: $(VENV)/.installed

$(VENV)/.installed: requirements.txt
	$(PYTHON) -m venv $(VENV)
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt
	$(ANSIBLE_GALAXY) collection install kubernetes.core
	touch $@

argo: venv
	cd $(ANSIBLE_DIR) && \
	$(ANSIBLE_PLAYBOOK) \
		-i inventory/hosts.yml \
		-e ansible_python_interpreter=$(VENV)/bin/python \
		playbooks/argocd.yml

argo-clean: venv
	cd $(ANSIBLE_DIR) && \
	$(ANSIBLE_PLAYBOOK) \
		-i inventory/hosts.yml \
		-e ansible_python_interpreter=$(VENV)/bin/python \
		playbooks/argocd-delete.yml

argo-reinstall: argo-clean argo

clean:
	rm -rf $(VENV)

poller-push:
	docker build -t harshith21/tekton-poller:2.0.0 ./tekton-poller
	docker tag harshith21/tekton-poller:2.0.0 harshith21/tekton-poller:latest
	docker push harshith21/tekton-poller:2.0.0
	docker push harshith21/tekton-poller:latest
