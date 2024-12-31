helm create sarkari-naukari-chart

go to the template folder and remove everything

and copy all the k8s manifests files 

go to deployment.yaml in helm and update the image with variable using jinja2 templating
image: abhishekp084/sarkari-naukari:{{ .Values.image.tag }}

remove everything from values.yaml and updatee the file with image and tag details

helm install sarkari-naukari ./sarkari-naukari-chart

helm uninstall
helm uninstall sarkari-naukari
