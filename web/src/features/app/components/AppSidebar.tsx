import {
  ChartArea,
  Folders,
  KeyRound,
  Layers,
  Settings,
  Users,
} from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuBadge,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/Sidebar";

// Menu items.
const items = [
  {
    title: "Files",
    url: "/app/files",
    icon: Folders,
    comingSoon: false,
  },
  {
    title: "Tokens",
    url: "/app/tokens",
    icon: KeyRound,
    comingSoon: false,
  },
  {
    title: "Logs",
    url: "#",
    icon: Layers,
    comingSoon: true,
  },
  {
    title: "Team",
    url: "#",
    icon: Users,
    comingSoon: true,
  },
  {
    title: "Usage",
    url: "#",
    icon: ChartArea,
    comingSoon: true,
  },
  {
    title: "Settings",
    url: "#",
    icon: Settings,
    comingSoon: true,
  },
];

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Fivemanage Lite (0.2.8)</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                  {item.comingSoon && (
                    <SidebarMenuBadge>Coming soon</SidebarMenuBadge>
                  )}
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
