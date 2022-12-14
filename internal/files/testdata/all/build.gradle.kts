apply plugin: "com.vrsn.long.paths.here.arent.they"

buildscript {
	repositories {
		def propertyResolver = new PropertyResolver(project)
		mavenCentral()
		if (propertyResolver.getRequiredBooleanProp("LOCAL_MAVEN_REPO_ENABLED", false)) {
			maven { url = propertyResolver.getRequiredStringProp("LOCAL_MAVEN_REPO") }
		}
		maven { url = "https://oss.sonatype.org/content/repositories/snapshots/" }
		gradlePluginPortal()
	}

	dependencies {
		classpath(BuildLibs.VRSN_GRADLE_PLUGIN)
	}
}

version = "0.9.12"
