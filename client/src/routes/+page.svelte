<script>
	import ActionBar from "./components/ActionBar.svelte";
	import Console from "./components/Console.svelte";
	import Editor from "./components/Editor.svelte";
	import Navbar from "./components/Navbar.svelte";

</script>
<div class="flex w-full flex-col">
  <Navbar />
  <div class="flex flex-row">
    <div class="flex flex-col grow">
      <Editor />
      <!-- Add the three  -->
      <Console />
    </div>
    <div class="flex flex-col grow w-[400px] ">
      <ActionBar />
    </div>
  </div>
</div>