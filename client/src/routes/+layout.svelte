<script>
	import { page } from "$app/stores";
  import "../app.css";
	import { dataContext } from "../store";



  // evlauate if page.form is nill
  $: if($page.form){
    $dataContext.outputAnalysis = $page.form.result
    $dataContext.TableSymbols = $page.form.symbol
    $dataContext.TableErrors = $page.form.errors == null ? [] : $page.form.errors
    $dataContext.cstTree = $page.form.dot
  }



</script>

<slot />