angular.module('controllers', []);
angular.module('directives', []);

var dependencyList = ['ngMaterial', 'ngRoute', 'controllers', 'directives'];

var themeColors = {
    primary: 'red', accent: 'deep-orange',
    warn: 'orange', background: 'grey'
};

var routes = {
    homeRoute: {
	    templateUrl: 'views/home.html',
	    controller: 'homeController'
    }
};

angular.module('badgerloop-dashboard', dependencyList)

.config(function($mdThemingProvider, $routeProvider) {
    /*************************  Theme Settings  ******************************/
	$mdThemingProvider
	.theme('default')
	.primaryPalette(themeColors.primary)
	.accentPalette(themeColors.accent)
	.warnPalette(themeColors.warn)
	.backgroundPalette(themeColors.background);
	$mdThemingProvider
	.theme('dark')
	.primaryPalette(themeColors.primary).dark();
    /*************************************************************************/

	$routeProvider
	.when('/', {redirectTo: '/home'})
    /*****************************  Routes  **********************************/
	.when('/home', routes.homeRoute)
    /*************************************************************************/
	.otherwise({redirectTo: 'home'})
});

