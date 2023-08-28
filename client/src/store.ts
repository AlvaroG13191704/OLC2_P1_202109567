import { writable } from "svelte/store";
import type { Analyzer } from "./interface";



// create a writable store with an initial value
export const dataContext = writable<Analyzer>({
  editorText: "",
  outputAnalysis: "",
  cstTree: "",
  TableErrors: [],
  TableSymbols: [],
});


