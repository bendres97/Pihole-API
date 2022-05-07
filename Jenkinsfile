pipeline 
{
    environment 
    {
        registry = "bryanendres/pihole-api"
        registry_cred = "Dockerhub"
        docker_versioned = ''
        docker_latest = ''
        versioned_image = "$registry:0.$BUILD_NUMBER"
    }
    
    agent any
    
    stages 
    {
        stage('Git gittin')
        {
            steps
            {
                git 'https://github.com/bendres97/Pihole-API.git'
            }
        }
        stage('Build') 
        {
            steps 
            {
                script
                {
                    docker_versioned = docker.build versioned_image
                    docker_latest = docker.build registry + ":latest"
                }
            }
        }
        stage('Vulnerability Scan')
        {
            steps
            {
                sh 'curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b . v0.27.1'
                sh './trivy image --ignore-unfixed --exit-code 1 --severity CRITICAL $docker_versioned'
            }
        }
        stage('Push to Docker')
        {
            steps
            {
                script
                {
                    docker.withRegistry('',registry_cred)
                    {
                        docker_versioned.push()
                        docker_latest.push()
                    }
                }
            }
        }
        stage('Clean up build images') 
        {
            steps
            {
                sh "docker rmi $versioned_image"
            }
        }
    }
}
