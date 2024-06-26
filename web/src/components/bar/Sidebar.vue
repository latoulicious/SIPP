<script setup lang="ts">
import { useAppStore } from "@/stores/index";
import { onMounted, computed } from "vue";

const appStore = useAppStore();

const showManagement = computed(() => {
  const jwtToken = localStorage.getItem("jwtToken");
  if (!jwtToken) {
    return false;
  }

  const decodedToken = decodeJwt(jwtToken);
  return decodedToken && decodedToken.role === "Admin";
});

const showPenilaian = computed(() => {
  // Replace this with your own logic
  return true;
});

onMounted(() => {
  // You might perform other setup actions here if needed
});

// Function to decode JWT token
function decodeJwt(jwtToken: string) {
  try {
    const [, payloadBase64] = jwtToken.split(".");
    const payload = JSON.parse(atob(payloadBase64));
    return payload;
  } catch (error) {
    console.error("Error decoding JWT:", error);
    return null;
  }
}
</script>

<template>
  <VaSidebar
    v-model="appStore.isSidebarVisible"
    color="#AFC1EE"
    class="custom-sidebar"
  >
    <router-link to="/dashboard">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="dashboard" />
          <VaSidebarItemTitle>Dashboard</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/capaian">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Capaian Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/alur">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Alur Tujuan Pendidikan</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/modul">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="folder_shared" />
          <VaSidebarItemTitle>Modul Ajar</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <VaAccordion>
      <VaCollapse v-show="showPenilaian">
        <template #header="{ value: isCollapsed }">
          <VaSidebarItem>
            <VaSidebarItemContent>
              <VaIcon name="assessment" />
              <VaSidebarItemTitle>Penilaian</VaSidebarItemTitle>
              <VaSpacer />
              <VaIcon :name="isCollapsed ? 'va-arrow-up' : 'va-arrow-down'" />
            </VaSidebarItemContent>
          </VaSidebarItem>
        </template>
        <template #body>
          <router-link to="/kognitif">
            <VaSidebarItem>
              <VaSidebarItemContent>
                <VaIcon name="assessment" />
                <VaSidebarItemTitle>Asesemen Kognitif</VaSidebarItemTitle>
              </VaSidebarItemContent>
            </VaSidebarItem>
          </router-link>
          <router-link to="/formatif">
            <VaSidebarItem>
              <VaSidebarItemContent>
                <VaIcon name="assessment" />
                <VaSidebarItemTitle>Asesmen Formatif</VaSidebarItemTitle>
              </VaSidebarItemContent>
            </VaSidebarItem>
          </router-link>
          <router-link to="/sumatif">
            <VaSidebarItem>
              <VaSidebarItemContent>
                <VaIcon name="assessment" />
                <VaSidebarItemTitle>Asesmen Sumatif</VaSidebarItemTitle>
              </VaSidebarItemContent>
            </VaSidebarItem>
          </router-link>
          <router-link to="/pengayaan">
            <VaSidebarItem>
              <VaSidebarItemContent>
                <VaIcon name="assessment" />
                <VaSidebarItemTitle>Pengayaan</VaSidebarItemTitle>
              </VaSidebarItemContent>
            </VaSidebarItem>
          </router-link>
          <router-link to="/remedial">
            <VaSidebarItem>
              <VaSidebarItemContent>
                <VaIcon name="assessment" />
                <VaSidebarItemTitle>Remedial</VaSidebarItemTitle>
              </VaSidebarItemContent>
            </VaSidebarItem>
          </router-link>
        </template>
      </VaCollapse>
    </VaAccordion>
    <router-link to="/soal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assessment" />
          <VaSidebarItemTitle>Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <router-link to="/banksoal">
      <VaSidebarItem>
        <VaSidebarItemContent>
          <VaIcon name="assessment" />
          <VaSidebarItemTitle>Bank Soal</VaSidebarItemTitle>
        </VaSidebarItemContent>
      </VaSidebarItem>
    </router-link>
    <VaAccordion>
      <VaCollapse v-show="showManagement">
        <template #header="{ value: isCollapsed }">
          <VaSidebarItem>
            <VaSidebarItemContent>
              <VaIcon name="settings" />
              <VaSidebarItemTitle>Management</VaSidebarItemTitle>
              <VaSpacer />
              <VaIcon :name="isCollapsed ? 'va-arrow-up' : 'va-arrow-down'" />
            </VaSidebarItemContent>
          </VaSidebarItem>
        </template>
        <template #body>
          <VaSidebarItem>
            <router-link v-show="showManagement" to="/kelas">
              <VaSidebarItem>
                <VaSidebarItemContent>
                  <VaIcon name="business" />
                  <VaSidebarItemTitle>Kelas</VaSidebarItemTitle>
                </VaSidebarItemContent>
              </VaSidebarItem>
            </router-link>
            <router-link v-show="showManagement" to="/jurusan">
              <VaSidebarItem>
                <VaSidebarItemContent>
                  <VaIcon name="business" />
                  <VaSidebarItemTitle>Jurusan</VaSidebarItemTitle>
                </VaSidebarItemContent>
              </VaSidebarItem>
            </router-link>
            <router-link v-show="showManagement" to="/mapel">
              <VaSidebarItem>
                <VaSidebarItemContent>
                  <VaIcon name="library_books" />
                  <VaSidebarItemTitle>Mata Pelajaran</VaSidebarItemTitle>
                </VaSidebarItemContent>
              </VaSidebarItem>
            </router-link>
            <router-link v-show="showManagement" to="/tahun">
              <VaSidebarItem>
                <VaSidebarItemContent>
                  <VaIcon name="date_range" />
                  <VaSidebarItemTitle>Tahun Ajaran</VaSidebarItemTitle>
                </VaSidebarItemContent>
              </VaSidebarItem>
            </router-link>
            <router-link v-show="showManagement" to="/user">
              <VaSidebarItem>
                <VaSidebarItemContent>
                  <VaIcon name="people" />
                  <VaSidebarItemTitle>User</VaSidebarItemTitle>
                </VaSidebarItemContent>
              </VaSidebarItem>
            </router-link>
          </VaSidebarItem>
        </template>
      </VaCollapse>
    </VaAccordion>
  </VaSidebar>
</template>

<style scoped>
.custom-sidebar {
  height: 100%;
  display: flex;
  flex-direction: column;
  overflow-y: auto; /* Enable scrolling if the content exceeds the sidebar height */
}

.custom-sidebar .VaSidebarItems {
  flex-grow: 1;
}

.custom-sidebar::-webkit-scrollbar {
  width: 12px; /* Width of the scrollbar */
}

.custom-sidebar::-webkit-scrollbar-thumb {
  background-color: #4a90e2; /* Color of the scrollbar thumb */
  border-radius: 6px; /* Border radius of the thumb */
}

.custom-sidebar::-webkit-scrollbar-track {
  background-color: #d3d3d3; /* Color of the scrollbar track */
  border-radius: 6px; /* Border radius of the track */
}

/* Hide the default scrollbar on Firefox */
.custom-sidebar {
  scrollbar-width: thin;
  scrollbar-color: #4a90e2 #d3d3d3;
}
</style>
