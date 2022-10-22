# Demo

For the demo we will run the [mesh-security](https://github.com/CosmWasm/mesh-security) tests against a setup
creted with kubeshuttle.

1. Clone 2 repos, mesh-security clone and kubeshuttle
   ```bash
   git clone git@github.com:Anmol1696/kubeshuttle.git
   
   # Mesh security repo, containing the tests
   git clone git@github.com:Anmol1696/mesh-security.git
   cd mesh-security
   git checkout anmol/k8s
   ```
   Note: We used the forked mesh-security repo since there are minor changes made to the testing repo. We will remove
   this dependency soon
2. If you dont have access to a k8s cluster, please follow the steps in [README](https://github.com/Anmol1696/kubeshuttle#setup-local-k8s-cluster-optional)
3. Optionally if you are resource constrained on your k8s cluster, you checkout kubeshuttle to a local branch, which will
   spin up less resources. If you have a beefy machine, no need... try it out
   ```bash
   cd kubeshuttle
   git checkout anmol/local
   ```
   Note: In this branch, we just change the `values.yaml` files to spin up only 2 chains, 1 relayer and explorer
4. Create the setup with kubeshuttle and port-forward ports
   ```bash
   cd kubeshuttle
   # Output all k8s yaml configurations
   make debug
   # Install cluster
   make install
   # Keep an eye on the pods that are spinning up
   k9s
   ```
   Note: If any of the `validator-*` node or relayer nodes is failing, just delete the pod and k8s will recreate one
   hopefully properly. This is a known issue, we are working on it. Use `kubectl delete pods <pod-name>`
5. Once the pods are running, run port-forward command
   ```bash
   cd kubeshuttle
   make port-forward-all
   ```
6. Now you can check the explorer locally at [http://localhost:8080](http://localhost:8080)
7. Run mesh-security tests against the system with
   ```bash
   cd mesh-security
   
   cd tests
   npm run tests
   ```
   Note: The tests could fail with error, `txn submitted but not on chain`. This is due to improper timeouts set on
   the clients used in the test scripts. Just re-run the tests again and hope the tests run :), while we fix the issue
8. All 3 tests shuold pass and you should be good to go.
