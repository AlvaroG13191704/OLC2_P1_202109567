import { analyzeAPI } from '$lib/API.js';
import { fail } from '@sveltejs/kit';
import type { Analyzer, Response } from '../interface.js';


export const actions = {
  analyze: async ({request}) => {
    const data = await request.formData();

    const code = data?.get('code')?.toString();

    // send code to server
    const response: Response = await analyzeAPI(code);

    // console.log(response);

    return response
  },
  openFile: async ({request}) => {
    const formData = Object.fromEntries(await request.formData());
    if (
      !(formData.fileToUpload as File).name ||
      (formData.fileToUpload as File).name === 'undefined'
    ) {
      return fail(400, {
        error: true,
        message: 'You must provide a file to upload'
      });
    }
    const { fileToUpload } = formData as { fileToUpload: File };

    // read file
    const fileContent = await fileToUpload.text();

    // return 
    return {
      fileContent
    }
  }
}