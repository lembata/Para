<script setup>
import { ref } from 'vue';

import AppMenuItem from './AppMenuItem.vue';

const model = ref([
	{
		label: 'Home',
		items: [{ label: 'Dashboard', icon: 'pi pi-fw pi-home', to: '/' }]
	},
	{
		label: 'Accounts',
		items: [
			{ label: 'List', icon: 'pi pi-fw pi-id-card', to: '/accounts' },
			{ label: 'Create', icon: 'pi pi-fw pi-check-square', to: '/accounts/add' },
			/*
			{ label: 'Float Label', icon: 'pi pi-fw pi-bookmark', to: '/uikit/floatlabel' },
			{ label: 'Invalid State', icon: 'pi pi-fw pi-exclamation-circle', to: '/uikit/invalidstate' },
			{ label: 'Button', icon: 'pi pi-fw pi-mobile', to: '/uikit/button', class: 'rotated-icon' },
			{ label: 'Table', icon: 'pi pi-fw pi-table', to: '/uikit/table' },
			{ label: 'List', icon: 'pi pi-fw pi-list', to: '/uikit/list' },
			{ label: 'Tree', icon: 'pi pi-fw pi-share-alt', to: '/uikit/tree' },
			{ label: 'Panel', icon: 'pi pi-fw pi-tablet', to: '/uikit/panel' },
			{ label: 'Overlay', icon: 'pi pi-fw pi-clone', to: '/uikit/overlay' },
			{ label: 'Media', icon: 'pi pi-fw pi-image', to: '/uikit/media' },
			{ label: 'Menu', icon: 'pi pi-fw pi-bars', to: '/uikit/menu', preventExact: true },
			{ label: 'Message', icon: 'pi pi-fw pi-comment', to: '/uikit/message' },
			{ label: 'File', icon: 'pi pi-fw pi-file', to: '/uikit/file' },
			{ label: 'Chart', icon: 'pi pi-fw pi-chart-bar', to: '/uikit/charts' },
			{ label: 'Misc', icon: 'pi pi-fw pi-circle', to: '/uikit/misc' }
			*/
		]
	},
	/*
	{
		label: 'Prime Blocks',
		items: [
			{ label: 'Free Blocks', icon: 'pi pi-fw pi-eye', to: '/blocks', badge: 'NEW' },
			{ label: 'All Blocks', icon: 'pi pi-fw pi-globe', url: 'https://www.primefaces.org/primeblocks-vue', target: '_blank' }
		]
	},
	{
		label: 'Utilities',
		items: [
			{ label: 'PrimeIcons', icon: 'pi pi-fw pi-prime', to: '/utilities/icons' },
			{ label: 'PrimeFlex', icon: 'pi pi-fw pi-desktop', url: 'https://www.primefaces.org/primeflex/', target: '_blank' }
		]
	},
	{
		label: 'Pages',
		icon: 'pi pi-fw pi-briefcase',
		to: '/pages',
		items: [
			{
				label: 'Landing',
				icon: 'pi pi-fw pi-globe',
				to: '/landing'
			},
			{
				label: 'Auth',
				icon: 'pi pi-fw pi-user',
				items: [
					{
						label: 'Login',
						icon: 'pi pi-fw pi-sign-in',
						to: '/auth/login'
					},
					{
						label: 'Error',
						icon: 'pi pi-fw pi-times-circle',
						to: '/auth/error'
					},
					{
						label: 'Access Denied',
						icon: 'pi pi-fw pi-lock',
						to: '/auth/access'
					}
				]
			},
			{
				label: 'Crud',
				icon: 'pi pi-fw pi-pencil',
				to: '/pages/crud'
			},
			{
				label: 'Timeline',
				icon: 'pi pi-fw pi-calendar',
				to: '/pages/timeline'
			},
			{
				label: 'Not Found',
				icon: 'pi pi-fw pi-exclamation-circle',
				to: '/pages/notfound'
			},
			{
				label: 'Empty',
				icon: 'pi pi-fw pi-circle-off',
				to: '/pages/empty'
			}
		]
	},
	{
		label: 'Hierarchy',
		items: [
			{
				label: 'Submenu 1',
				icon: 'pi pi-fw pi-bookmark',
				items: [
					{
						label: 'Submenu 1.1',
						icon: 'pi pi-fw pi-bookmark',
						items: [
							{ label: 'Submenu 1.1.1', icon: 'pi pi-fw pi-bookmark' },
							{ label: 'Submenu 1.1.2', icon: 'pi pi-fw pi-bookmark' },
							{ label: 'Submenu 1.1.3', icon: 'pi pi-fw pi-bookmark' }
						]
					},
					{
						label: 'Submenu 1.2',
						icon: 'pi pi-fw pi-bookmark',
						items: [{ label: 'Submenu 1.2.1', icon: 'pi pi-fw pi-bookmark' }]
					}
				]
			},
			{
				label: 'Submenu 2',
				icon: 'pi pi-fw pi-bookmark',
				items: [
					{
						label: 'Submenu 2.1',
						icon: 'pi pi-fw pi-bookmark',
						items: [
							{ label: 'Submenu 2.1.1', icon: 'pi pi-fw pi-bookmark' },
							{ label: 'Submenu 2.1.2', icon: 'pi pi-fw pi-bookmark' }
						]
					},
					{
						label: 'Submenu 2.2',
						icon: 'pi pi-fw pi-bookmark',
						items: [{ label: 'Submenu 2.2.1', icon: 'pi pi-fw pi-bookmark' }]
					}
				]
			}
		]
	},
	{
		label: 'Get Started',
		items: [
			{
				label: 'Documentation',
				icon: 'pi pi-fw pi-question',
				to: '/documentation'
			},
			{
				label: 'Figma',
				url: 'https://www.dropbox.com/scl/fi/bhfwymnk8wu0g5530ceas/sakai-2023.fig?rlkey=u0c8n6xgn44db9t4zkd1brr3l&dl=0',
				icon: 'pi pi-fw pi-pencil',
				target: '_blank'
			},
			{
				label: 'View Source',
				icon: 'pi pi-fw pi-search',
				url: 'https://github.com/primefaces/sakai-vue',
				target: '_blank'
			},
			{
				label: 'Nuxt Version',
				url: 'https://github.com/primefaces/sakai-nuxt',
				icon: 'pi pi-fw pi-star'
			}
		]
	}
	*/
]);
</script>

<template>
	<ul class="layout-menu">
		<template v-for="(item, i) in model" :key="item">
			<app-menu-item v-if="!item.separator" :item="item" :index="i"></app-menu-item>
			<li v-if="item.separator" class="menu-separator"></li>
		</template>
	</ul>
</template>
<!--
            <a href="https://www.primefaces.org/primeblocks-vue/#/" target="_blank">
                <img src="/layout/images/banner-primeblocks.png" alt="Prime Blocks" class="w-full mt-3" />
            </a>
-->

<style>
.layout-menu {
	@apply 
		m-0 p-0
		list-none
	;

	a {
		@apply select-none;

		&.active-menuitem {
			>.layout-submenu-toggler {
				@apply rotate-180;
			}
		}
	}

	li.active-menuitem {
		>a {
			.layout-submenu-toggler {
				@apply rotate-180;
			}
		}
	}

	ul {
		@apply m-0 p-0 list-none focus:ring-2;

		a {
			@apply flex items-center relative
				outline-none font-bold
				text-gray-100 
				hover:text-gray-300 
				cursor-pointer 
				py-3 px-4 rounded transition-colors transition-shadow duration-200;

			.layout-menuitem-icon {
				@apply mr-2 text-gray-200;
				
			}

			.layout-submenu-toggler {
				font-size: 75%;
				@apply ml-auto transition-transform;
			}

			&.active-route {
				@apply font-bold text-primary-500;
			}

		}

		ul {
			overflow: hidden;
			border-radius: 2px;

			li {
				a {
					@apply ml-4;
				}

				li {
					a {
						@apply ml-8;
					}

					li {
						a {
							@apply ml-10;
						}

						li {
							a {
								@apply ml-12;
							}

							li {
								a {
									@apply ml-14;
								}

								li {
									a {
										@apply ml-16;
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
</style>
