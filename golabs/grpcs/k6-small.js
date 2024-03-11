import grpc from "k6/net/grpc";
//import { check, sleep } from "k6";
import exec from "k6/execution";
import { Rate } from "k6/metrics";
const client = new grpc.Client();
client.load(["./"], "dto.proto", "service.proto");

// Create a rate for the number of requests
const grpc_reqs = new Rate("grpc_reqs");

export const options = {
  thresholds: {
    grpc_req_duration: [
      {
        threshold: "p(95)<=100", // Set to a very low value to test thresholds
        abortOnFail: true,
        delayAbortEval: "1s",
      },
      {
        threshold: "p(99.999)<=300",
      },
    ],
  },
};

export default () => {
  if (exec.vu.iterationInScenario == 0) {
  client.connect("0.0.0.0:8080", {
    plaintext: true,
    // reflect: true,
  });
    }


    const data = {
      name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    };

    client.invoke("main.Test/Test", data);
    // Add one to the counter for each request
    // grpc_reqs.add(response.status === grpc.StatusOK);
    // check(response, {
    //   "status is OK": (r) => r && r.status === grpc.StatusOK,
    // });
  

  //   console.log(JSON.stringify(response.message));

  //   client.close();
  //   sleep(1);
};
