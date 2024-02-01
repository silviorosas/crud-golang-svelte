<!-- PersonaTable.svelte -->
<script>
    let personas = [];
  
    // Función para cargar los datos de las personas desde el servidor Go
    async function fetchPersonas() {
      const response = await fetch('http://localhost:9000/personas'); // Reemplaza la URL con la dirección de tu servidor Go
      const data = await response.json();
      personas = data;
      console.log("Personas desde Go:", data)
    }
   
    // Cargar los datos al montar el componente
    fetchPersonas();
  </script>
  
  <h3>Lista de Personas</h3>
  <table class="table table-dark table-hover w-75">
    <thead>
      <tr>
        <th>ID</th>
        <th>Nombre</th>
        <th>Apellido</th>
        <th>Email</th>
        <th>Teléfono</th>
      </tr>
    </thead>
    <tbody>
      {#each personas as persona}
        <tr>
          <td>{persona.id}</td>
          <td>{persona.nombre}</td>
          <td>{persona.apellido}</td>
          <td>{persona.email}</td>
          <td>{persona.telefono}</td>
        </tr>
      {/each}
    </tbody>
  </table>

  <style>
   
  </style>
  