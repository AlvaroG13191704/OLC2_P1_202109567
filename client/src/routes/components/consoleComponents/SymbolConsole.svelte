<script>
	import { dataContext } from "../../../store";


 
  /** @type {any[]} */
  $: symbols = $dataContext.TableSymbols;


</script>

<div class="bg-yellow-100 p-4 rounded mt-2 h-full">
  <h3 class="text-yellow-700 font-semibold mb-2">Tabla de simbolos:</h3>
  <table class="w-full">
    <thead>
      <tr class=" text-yellow-700">
        <th class="px-4 py-2">ID</th>
        <th class="px-4 py-2">Tipo Símbolo</th>
        <th class="px-4 py-2">Tipo dato</th>
        <th class="px-4 py-2">Línea</th>
        <th class="px-4 py-2">Columna</th>
      </tr>
    </thead>
    <tbody>
      {#each symbols as symbol (symbol)}
        <tr class="bg-white border-b border-gray-300 text-center">
          <td class="px-4 py-2">{symbol.Id}</td>
          <td class="px-4 py-2">{symbol.TypeSymbol}</td>
          <td class="px-4 py-2">{symbol.TypeData}</td>
          <td class="px-4 py-2">{symbol.Line}</td>
          <td class="px-4 py-2">{symbol.Column}</td>
        </tr>
      {/each}
    </tbody>
  </table>  
</div>
