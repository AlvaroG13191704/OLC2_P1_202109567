<script>
// @ts-nocheck

	import { enhance } from '$app/forms';

  // get the store
  import { dataContext } from '../../store';



</script>



<div class="flex h-full ">
  <!-- Sidebar -->
  <form class="w-full bg-gray-900 text-white" method="POST" action="?/analyze" use:enhance  >
    <input type="hidden" name="code" value={$dataContext.editorText} />
    <div class="py-6 px-4">
      <h2 class="text-2xl text-center font-semibold">Acciones</h2>
    </div>
    <ul class="space-y-16 m-4 ">
      <li>
        <button  class="w-full py-4 px-6 text-lg font-semibold rounded-lg bg-green-500 hover:bg-green-800">
          EJECUTAR
        </button>
      </li>
      <li>
        <button class="w-full py-4 px-6 text-lg font-semibold rounded-lg bg-gray-700 hover:bg-gray-600">
          MOSTRAR REPORTES
        </button>
      </li>
    </ul>
  </form>
</div>