<script>
	import ActionBar from "./components/ActionBar.svelte";
	import Editor from "./components/Editor.svelte";
	import Navbar from "./components/Navbar.svelte";

</script>
<div class="flex min-h-full w-full flex-col">
  <Navbar />
  <div class="flex flex-row h-full">
    <div class="flex flex-col grow">
      <Editor />
      <!-- Add the  -->
    </div>
    <div class="flex flex-col grow ">
      <ActionBar />
    </div>
  </div>
</div>