.PHONY: test
test:
	../scripts/cleanup.sh                                 # cleanup
	../scripts/setup.sh                                   # create namespace, CRDs, composition & facade CRs
	kubectl apply -f ../manifests/test1/plan1.yaml    # create plan CR
	kubectl patch cloudsqls cloudsql-sample -n facade --type=merge --subresource status --patch 'status: {expanders: [{"name": "block2", "jsonValues": "{\"identity\": { \"email\": \"foobar@myproj.google.com\"} }"},]}'
	kubectl apply -f ./inline-jinja-job.yaml              # run Job to expand 
	sleep 30
	echo "Status after wait:"                             # verify facade CR status after expansion
	../scripts/check_plan.sh cloudsqls-cloudsql-sample facade ../manifests/test1/cloudsql1.expanded
	../scripts/cleanup.sh                                 # cleanup

