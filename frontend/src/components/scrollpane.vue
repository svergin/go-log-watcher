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
        
        const reader =(await res.body?.getReader());
        //this.loglines = (await res.json()).lines
        reader?.read().then(({done, value }) => {
          // if(done) {
          //   return res
          // }
          this.loglines = String.fromCharCode.apply(null, value)
          
        })
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
    <textarea v-model="loglines" cols="50" rows="25"/>
    
  </div>
</template>