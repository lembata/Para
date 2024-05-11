<script setup>
import { computed, watch, ref } from 'vue';
import AppTopbar from './AppTopbar.vue';
import AppFooter from './AppFooter.vue';
import AppSidebar from './AppSidebar.vue';
//import AppConfig from './AppConfig.vue';
import { useLayout } from '@/layout/composables/layout';

const { layoutConfig, layoutState, isSidebarActive } = useLayout();

const outsideClickListener = ref(null);

watch(isSidebarActive, (newVal) => {
    if (newVal) {
        bindOutsideClickListener();
    } else {
        unbindOutsideClickListener();
    }
});

const containerClass = computed(() => {
    return {
        'layout-theme-light': layoutConfig.darkTheme.value === 'light',
        'layout-theme-dark': layoutConfig.darkTheme.value === 'dark',
        'layout-overlay': layoutConfig.menuMode.value === 'overlay',
        'layout-static': false,//layouttext-Config.menuMode.value === 'static',
        'layout-static-inactive': layoutState.staticMenuDesktopInactive.value && layoutConfig.menuMode.value === 'static',
        'layout-overlay-active': layoutState.overlayMenuActive.value,
        'layout-mobile-active': layoutState.staticMenuMobileActive.value,
        'p-ripple-disabled': layoutConfig.ripple.value === false
    };
});

const bindOutsideClickListener = () => {
    if (!outsideClickListener.value) {
        outsideClickListener.value = (event) => {
            if (isOutsideClicked(event)) {
                layoutState.overlayMenuActive.value = false;
                layoutState.staticMenuMobileActive.value = false;
                layoutState.menuHoverActive.value = false;
            }
        };
        document.addEventListener('click', outsideClickListener.value);
    }
};

const unbindOutsideClickListener = () => {
    if (outsideClickListener.value) {
        document.removeEventListener('click', outsideClickListener);
        outsideClickListener.value = null;
    }
};

const isOutsideClicked = (event) => {
    const sidebarEl = document.querySelector('.layout-sidebar');
    const topbarEl = document.querySelector('.layout-menu-button');

    return !(sidebarEl.isSameNode(event.target) || sidebarEl.contains(event.target) || topbarEl.isSameNode(event.target) || topbarEl.contains(event.target));
};
</script>

<template>
    <div class="layout-wrapper layout-static" :class="containerClass">
        <app-topbar></app-topbar>
        <div class="layout-sidebar">
            <app-sidebar></app-sidebar>
        </div>
        <div class="layout-main-container">
            <div class="layout-main">
                <router-view></router-view>
            </div>
        </div>
        <app-config></app-config>
        <div class="layout-mask"></div>
    </div>
    <Toast />
</template>

<style>
.layout-sidebar {
    @apply fixed 
        w-72 
        top-28 left-4 
        py-2 px-6 
		bg-surface-800
        text-gray-100
        transition-transform 
        ease-in-out 
        duration-300 
        select-none 
        overflow-y-auto 
        rounded-sm 
        shadow;

    height: calc(100vh - 9rem);
    z-index: 999;

    /*
    transition: transform $transitionDuration, left $transitionDuration;
    box-shadow: 0px 3px 5px rgba(0, 0, 0, 0.02), 0px 0px 2px rgba(0, 0, 0, 0.05), 0px 1px 4px rgba(0, 0, 0, 0.08);
    border-radius: $borderRadius;
     */
}


.ml-auto-important {
    margin-left: auto !important;
}

.mr-auto-important {
    margin-right: auto !important;
}

.layout-main,
.landing-wrapper {
    @apply w-full
        ml-auto-important 
        mr-auto-important;
}

.translate-x-full-reverse {
    transform: translateX(-100%);
}

.fade-in {
    animation: fadein 0.5s; /*time?*/
}

.layout-wrapper {
    &.layout-static {
        .layout-main-container {
            @apply ml-72;
        }
    }

    @apply w-screen h-screen;

}

/*
@screen md {
    .layout-wrapper {
        &.layout-overlay {
            .layout-main-container {
                @apply ml-0 pl-8;

                .layout-sidebar {
            @apply 
                translate-x-full-reverse
                left-0 top-0
                h-screen
                rounded-tl-none rounded-bl-none;
                }

                &.layout-overlay-active {
                    .layout-sidebar {
                        @apply translate-x-0;
                    }
                }
            }
        }

        &.layout-static {
            .layout-main-container {
                @apply ml-72;
            }

            &.layout-static-inactive {
                .layout-sidebar {
                    @apply left-0 translate-x-full-reverse;
                }

                .layout-main-container {
                    @apply ml-0 pl-8;
                }
            }
        }

        .layout-mask {
            @apply hidden;
        }
    }
}
    */


@screen md {
    .blocked-scroll {
        @apply overflow-hidden;
    }

    .layout-wrapper {
        .layout-main-container {
            @apply ml-0 pl-8;
        }

        .layout-sidebar {
            @apply 
                translate-x-full-reverse
                left-0 top-0
                h-screen
                rounded-tl-none rounded-bl-none;
        }

        .layout-mask {
            @apply hidden
                fixed
                top-0 left-0
                w-full h-full
                bg-opacity-20
                bg-gray-800;
                
            z-index: 998;
        }

        &.layout-mobile-active {
            .layout-sidebar {
                @apply translate-x-0;
            }

            .layout-mask {
                @apply block fade-in;
            }
        }

        &.layout-static {
            .layout-main-container {
                @apply ml-0;
            }
        }
    }
}
</style>
