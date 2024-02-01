<!-- AddPersonaForm.svelte -->
<script>
  import { onMount, createEventDispatcher } from 'svelte';

  let persona = { nombre: '', apellido: '', email: '', telefono: '' };
  let errorMessage = '';

  const dispatch = createEventDispatcher();

  function handleSubmit(event) {
    event.preventDefault();

    // Enviar datos al endpoint Go
    fetch('http://localhost:9000/personas', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(persona)
    })
    .then(response => {
      if (!response.ok) {
        return response.json().then(data => {
          throw new Error(data.error || 'Error al enviar los datos');
        });
      }
      return response.json();
    })
    .then(data => {
      // Enviar datos al componente padre
      dispatch('personaAgregada', { persona: data });
      // Reiniciar el formulario
      persona = { nombre: '', apellido: '', email: '', telefono: '' };
      errorMessage = ''; // Limpiar el mensaje de error
    })
    .catch(error => {
      console.error('Error al enviar los datos:', error.message);
      errorMessage = error.message;
    });
  } 
</script>

<div class="container w-50">
  <h3>Form Agregar Persona</h3>
  <form on:submit={handleSubmit}> 
      <div class="mb-3">
          <label for="nombre" class="form-label">Nombre</label>
          <input type="text" class="form-control" id="nombre" bind:value={persona.nombre}>
      </div>
      <div class="mb-3">
          <label for="apellido" class="form-label">Apellido</label>
          <input type="text" class="form-control" id="apellido" bind:value={persona.apellido}>
      </div>
      <div class="mb-3">
          <label for="email" class="form-label">Email</label>
          <input type="text" class="form-control" id="email" bind:value={persona.email}>
          {#if errorMessage && errorMessage.includes('correo electrónico ya está registrado')}
              <div class="alert alert-danger" role="alert">Este correo electrónico ya está registrado</div>
          {/if}
      </div>
      <div class="mb-3">
          <label for="telefono" class="form-label">Teléfono</label>
          <input type="text" class="form-control" id="telefono" bind:value={persona.telefono}>
      </div>
      <button type="submit" class="btn btn-primary">Agregar Persona</button>
  </form>
</div>
