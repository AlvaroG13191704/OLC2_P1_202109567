<script lang="ts">
	import { dataContext } from '../../store';
  
  import loader from '@monaco-editor/loader';
  import { onDestroy, onMount } from 'svelte';
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let monaco: typeof Monaco;
  let editorContainer: HTMLElement;


  onMount(async () => {

      // Remove the next two lines to load the monaco editor from a CDN
      // see https://www.npmjs.com/package/@monaco-editor/loader#config
      const monacoEditor = await import('monaco-editor');
      loader.config({ monaco: monacoEditor.default });

      monaco = await loader.init();

      // Your monaco instance is ready, let's display some code!
      const editor = monaco.editor.create(editorContainer, {
        automaticLayout: true,
        theme: 'vs-dark',
        autoClosingBrackets: 'always',
        folding: true,
        lineNumbersMinChars: 3,
        scrollBeyondLastLine: false,
        minimap: {
          enabled: false,
        },
        
      });
      const model = monaco.editor.createModel(
          "print(\"Hello World from T-swift\")",
          'swift',
          // Give monaco a hint which syntax highlighting to use
          monaco.Uri.file('sample.swift')
      );
      
      editor.setModel(model);

      // Set the editor value
      editor.onDidChangeModelContent(() => {
        $dataContext.editorText = editor.getValue();
      });
  });

  onDestroy(() => {
      monaco?.editor.getModels().forEach((model) => model.dispose());
  });
</script>

<div>
  <div class="container" bind:this={editorContainer} />
</div>

<style>
  .container {
      width: 100%;
      height: 500px;
  }
</style>