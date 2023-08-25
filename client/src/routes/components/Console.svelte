
<script>
	import { dataContext } from "../../store";
	import ErrorConsole from "./consoleComponents/ErrorConsole.svelte";
	import SymbolConsole from "./consoleComponents/SymbolConsole.svelte";



  let activeConsole = 'main' // 'symbols' , 'errors'


</script>



<div class="bg-gray-900 h-[360px] text-lg p-2 "> 
  <div >
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'main'}>Consola</button>
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'symbols'}>Simbolos</button>
    <button class="bg-gray-700 text-gray-200 px-2 py-1 rounded-md hover:bg-gray-500" on:click={() => activeConsole = 'errors'}>Errores</button>
  </div>
  
  <div class="border-b-2 border-gray-700 h-[300px] m-2">
  {#if activeConsole === 'main'}
    <textarea class="bg-slate-300 text-black w-full h-full p-3 font-bold" disabled bind:value={$dataContext.outputAnalysis} />
  {:else if activeConsole === 'symbols'}
    <SymbolConsole  />
  {:else if activeConsole === 'errors'}
    <ErrorConsole  />
  {/if}        
  </div>
  

</div>