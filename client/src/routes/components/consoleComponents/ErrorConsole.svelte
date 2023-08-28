<script>
	import { dataContext } from "../../../store";

  /** @type {any[]} */
  $: errors = $dataContext.TableErrors == null ? [] : $dataContext.TableErrors;
  
</script>
<div class="bg-red-100 p-4 h-full mt-2 rounded">
  <h3 class="text-red-700 font-semibold mb-2">Errores:</h3>
  <table class="w-full">
    <thead>
      <tr class="text-red-700">
        <th class="px-4 py-2">No.</th>
        <th class="px-4 py-2">Mensaje</th>
        <th class="px-4 py-2">Tipo</th>
        <th class="px-4 py-2">Línea</th>
        <th class="px-4 py-2">Columna</th>
      </tr>
    </thead>
    <tbody>
      {#each errors as error , i}
        <tr class="bg-white border-b border-gray-300 text-center">
          <td class="px-4 py-2">{i+1}</td>
          <td class="px-4 py-2">{error.Msg}</td>
          <td class="px-4 py-2">{error.Type}</td>
          <td class="px-4 py-2">{error.Line}</td>
          <td class="px-4 py-2">{error.Column}</td>
        </tr>
      {/each}
    </tbody>
  </table> 
</div>
