// auth.js

import axios from "axios"; // Make sure to install and import axios

export default async function (to, from, next) {
  const jwtToken = localStorage.getItem("jwtToken");
  console.log("Retrieved token:", jwtToken);

  // Check if the user is authenticated (you might want to add more checks)
  if (!jwtToken) {
    // Redirect to the login page if not authenticated
    next({ name: "login" }); // Use the name of the route
  } else {
    // Fetch user role from the API endpoint
    const userRole = await fetchUserRole(jwtToken);

    // Check if the user has the required role to access the route
    if (to.name === "user" && userRole !== "admin") {
      // If the user is not an admin, redirect to a different route or show an error
      next({ name: "dashboard" }); // Redirect to the dashboard for example
    } else {
      // Proceed to the next route if authenticated and has the required role
      next();
    }
  }
}

async function fetchUserRole(jwtToken) {
  try {
    const response = await axios.get(
      "http://localhost:3000/api/get-user-role",
      {
        headers: {
          Authorization: `Bearer ${jwtToken}`,
          "Content-Type": "application/json",
        },
      },
    );

    // Log the entire response for debugging
    console.log("Response from /api/get-user-role:", response);

    // Return the user role from the response
    return response.data.role;
  } catch (error) {
    console.error("Error fetching user role:", error);
    return null;
  }
}
