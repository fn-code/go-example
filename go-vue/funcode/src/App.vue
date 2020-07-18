<template>

  <v-app>
    <TemplateLayout/>
  </v-app>


</template>
<script>
import TemplateLayout  from "./components/TemplateLayout";

export default {
  name: 'App',
  components: {
    TemplateLayout
  }
}
</script>

<style>
*{
  font-family: 'Roboto', sans-serif;
}
</style>
