import { analyzeAPI } from '$lib/API.js';
import type { Analyzer, Response } from '../interface.js';


export const actions = {
  analyze: async ({request}) => {
    const data = await request.formData();

    const code = data?.get('code')?.toString();

    // send code to server
    const response: Response = await analyzeAPI(code);

    console.log(response);

    return response
  }
}