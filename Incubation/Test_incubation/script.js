document.addEventListener("DOMContentLoaded", () => {
    const input = document.getElementById("service-input");
    const suggestions = document.getElementById("suggestions");
  
    // Show suggestions when the input is focused
    input.addEventListener("focus", () => {
      suggestions.classList.remove("hidden");
    });
  
    // Hide suggestions when clicking outside the input or suggestions
    document.addEventListener("click", (event) => {
      if (!input.contains(event.target) && !suggestions.contains(event.target)) {
        suggestions.classList.add("hidden");
      }
    });
  
    // Handle suggestion click
    suggestions.addEventListener("click", (event) => {
      if (event.target.tagName === "LI") {
        input.value = event.target.textContent;
        suggestions.classList.add("hidden");
      }
    });
  
    // Hide suggestions when typing
    input.addEventListener("input", () => {
      suggestions.classList.add("hidden");
    });
  });
  

  // redirection page nounou

  // Récupérer l'élément select
  const categorySelect = document.getElementById("category-select");

  // Ajouter un écouteur d'événement pour le changement d'option
  categorySelect.addEventListener("change", function() {
    const selectedValue = categorySelect.value;

    // Vérifier si l'option sélectionnée est "nounou"
    if (selectedValue === "nounou") {
      // Rediriger vers la page des nounous
      window.location.href = "page_nounou.html"; // Remplacez par le lien réel
    }
    // Vous pouvez ajouter d'autres redirections pour d'autres options
    else if (selectedValue === "creches") {
      window.location.href = "page_creches.html"; // Remplacez par le lien réel
    }
    else if (selectedValue === "gardes") {
      window.location.href = "page_gardes.html"; // Remplacez par le lien réel
    }
  });