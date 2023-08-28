<script>
	import { enhance } from "$app/forms";
  import { page } from "$app/stores";
	import { dataContext } from "../../store";

  $: if ($page.form) {
    $dataContext.editorText = $page.form.fileContent;
  }

  async function save() {
    // open the windows to save the file
    console.log($dataContext.editorText);
    const blob = new Blob([$dataContext.editorText], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = "file.swift";
    link.click();
    URL.revokeObjectURL(url);
  }

  async function saveAs() {
    // open the windows to save the file
    const blob = new Blob([$dataContext.editorText], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = "file.swift";
    link.click();
    URL.revokeObjectURL(url);
  }

</script>

<header class="bg-gray-900">
  <nav class="flex items-center justify-start  p-6 lg:px-8 text-white" aria-label="Global">
    <form method="POST"  enctype="multipart/form-data">
      <!-- hidden the editorText -->
      <input type="hidden" name="editorText" bind:value={$dataContext.editorText} />
      <input type="file" id="file" name="fileToUpload" class="mr-10 py-2 px-6 text-sm font-semibold rounded-lg bg-gray-700 hover:bg-gray-600" required  />
      <button formaction="?/openFile" class="mr-10 py-2 px-6 text-lg font-semibold rounded-lg bg-gray-700 hover:bg-gray-600">
        Subir
      </button>
    </form>
    <button on:click={save} class="mr-10 py-2 px-6 text-lg font-semibold rounded-lg bg-gray-700 hover:bg-gray-600">
      Guardar
    </button>
    <button on:click={saveAs} class="py-2 px-6 text-lg font-semibold rounded-lg bg-gray-700 hover:bg-gray-600">
      Guardar como
    </button>
  </nav>
</header>