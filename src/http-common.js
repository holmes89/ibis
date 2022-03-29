import Auth from "@aws-amplify/auth";
import axios from "axios";

const apiEndpoint =
  process.env.REACT_APP_API_URL === undefined
    ? ""
    : process.env.REACT_APP_API_URL;
const getToken = async () => {
  const session = await Auth.currentSession();
  const idToken = session.getIdToken();
  return idToken.getJwtToken();
};

const client = axios.create({
  baseURL: apiEndpoint + "/",
  headers: {
    "Content-type": "application/json",
  },
});

client.interceptors.request.use(
  async (config) => {
    const idToken = await getToken();
    return {
      ...config,
      headers: {
        Authorization: `Bearer ${idToken}`,
      },
    };
  },
  (error) => Promise.reject(error)
);

export default client;
