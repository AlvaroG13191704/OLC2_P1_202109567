<script>
// @ts-nocheck

	import { enhance } from '$app/forms';
	import { onMount } from 'svelte';

  // get the store
  import { dataContext } from '../../store';



  let modal = false;

  const setModal = () => {
    if (modal) {
      modal = false;
    } else {
      modal = true;
    }
  }
  

</script>



<div class="flex h-full ">
  <!-- Sidebar -->
  <div class="w-full bg-gray-900 text-white"   >
    <div class="py-6 px-4">
      <h2 class="text-2xl text-center font-semibold">Acciones</h2>
    </div>
    <ul class="space-y-16 m-4 ">
      <form method="POST" action="?/analyze" use:enhance>
        <li>
          <input type="hidden" name="code" value={$dataContext.editorText} />
          <button  class="w-full py-4 px-6 text-lg font-semibold rounded-lg bg-green-500 hover:bg-green-800">
            EJECUTAR
          </button>
        </li>
      </form>
      <li>
        <button  on:click={setModal} class="w-full py-4 px-6 text-lg font-semibold rounded-lg bg-gray-700 hover:bg-gray-600">
          MOSTRAR REPORTE CST
        </button>
      </li>
    </ul>
  </div>
</div>

{#if modal === true}
  // Create a modal to set an svg image  
  <div class="relative z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
  <div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>
  <div class="fixed inset-0 z-10 m-10">
    <div class="bg-white">
      {@html $dataContext.cstTree}
      <div class="mt-5 sm:mt-6">
        <button type="button" on:click={setModal} class="inline-flex w-full justify-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Regresar</button>
      </div>
    </div>
  </div>
</div>

{/if}
