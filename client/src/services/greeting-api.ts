import { Configuration } from '../api/greeting-api-client';
import { GreetingApi } from '../api/greeting-api-client/api';

const config = new Configuration({
  basePath: import.meta.env.VITE_API_BASE_URL,
});

export default new GreetingApi(config);
