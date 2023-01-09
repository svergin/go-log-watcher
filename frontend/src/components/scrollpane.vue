<script lang ="ts">
export default {
  data() {
    return {
      loglines: "Waiting..."
    }
  },
  methods: {
    async getAnswer() {
      
      try {
        const res = await fetch('http://localhost:8080/log/watch')
        this.loglines = (await res.json()).lines
      } catch (error) {
        this.loglines = 'Error! Could not reach the API. ' + error
      }
    }
  },
  mounted() {
    this.getAnswer()
  }
}
</script>
<template>
  <div class="logwelcome">
    <h1 class="green">Logfile</h1>
  </div>
  <div>
    <textarea  v-model="loglines" size="50"/>
    
  </div>
</template>