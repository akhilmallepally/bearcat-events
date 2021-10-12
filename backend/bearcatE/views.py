from django.shortcuts import render
from django.http import HttpResponse
from rest_framework import viewsets
from .models import EventModel
from .serializers import bearcatESerializer

# Create your views here.
def home(request):
    return HttpResponse("This is home")
class bearcatE_view(viewsets.ModelViewSet):
    serializer_class = bearcatESerializer
    queryset = EventModel.objects.all()
